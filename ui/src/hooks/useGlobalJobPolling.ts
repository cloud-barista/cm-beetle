'use client';

import { useEffect, useRef } from 'react';
import { useMigrationStore, calculateJobElapsedSeconds } from '../store/migrationStore';
import { beetleApi } from '../api/client';
import { RequestDetailsResult } from '../types/migration';

export function useGlobalJobPolling() {
  const {
    jobs,
    setJobs,
    objectStorageJobs,
    setObjectStorageJobs,
    dataJobs,
    setDataJobs
  } = useMigrationStore();

  const jobsRef = useRef(jobs);
  const objectStorageJobsRef = useRef(objectStorageJobs);
  const dataJobsRef = useRef(dataJobs);
  const lastDataPollAtRef = useRef<Record<string, number>>({});
  const lastInfraNodePollAtRef = useRef<Record<string, number>>({});

  useEffect(() => {
    jobsRef.current = jobs;
  }, [jobs]);

  useEffect(() => {
    objectStorageJobsRef.current = objectStorageJobs;
  }, [objectStorageJobs]);

  useEffect(() => {
    dataJobsRef.current = dataJobs;
  }, [dataJobs]);

  useEffect(() => {
    const interval = setInterval(async () => {
      const currentJobs = jobsRef.current;
      const currentOsJobs = objectStorageJobsRef.current;
      const currentDataJobs = dataJobsRef.current;

      const activeInfraJobs = currentJobs.filter((j) => !j.isSample && j.status === 'Handling');
      const activeNlbJobs = currentJobs.filter((j) => !j.isSample && j.nlbStatus === 'Provisioning' && j.nlbReqId);
      const activeOsJobs = currentOsJobs.filter((j) => !j.isSample && j.status === 'Handling');
      const activeDataJobs = currentDataJobs.filter((j) => !j.isSample && j.status === 'Handling');

      if (
        activeInfraJobs.length === 0 &&
        activeNlbJobs.length === 0 &&
        activeOsJobs.length === 0 &&
        activeDataJobs.length === 0
      ) {
        return;
      }

      // 1. Poll Active Infrastructure Migration Jobs & NLB Stage 2 Jobs
      if (activeInfraJobs.length > 0 || activeNlbJobs.length > 0) {
        const statusUpdates: Record<string, RequestDetailsResult> = {};
        const liveNodeUpdates: Record<string, any[]> = {};
        const nlbStatusUpdates: Record<string, RequestDetailsResult> = {};
        const nowMs = Date.now();

        for (const job of activeInfraJobs) {
          try {
            let res = await beetleApi.getRequestDetails(job.reqId, { brief: true });
            const isTerminalSuccess =
              res?.status === 'Success' || res?.status === 'Completed' || res?.status === 'Succeeded';
            if (isTerminalSuccess && res.responseData === undefined) {
              res = await beetleApi.getRequestDetails(job.reqId, { brief: false });
            }
            if (res && res.status) {
              statusUpdates[job.id] = res;
            }
            const infraStepActive =
              res?.status === 'Handling' &&
              Array.isArray(res.subRequests) &&
              res.subRequests.some((sr) => sr.resourceType === 'infra' && sr.method === 'POST');
            if (infraStepActive && job.nsId && job.infraId) {
              const lastNodePoll = lastInfraNodePollAtRef.current[job.id] || 0;
              if (nowMs - lastNodePoll >= 5000) {
                lastInfraNodePollAtRef.current[job.id] = nowMs;
                const liveDetail = await beetleApi.getMigratedInfraDetail(job.nsId, job.infraId);
                if (liveDetail && Array.isArray(liveDetail.node) && liveDetail.node.length > 0) {
                  liveNodeUpdates[job.id] = liveDetail.node;
                }
              }
            }
          } catch (err) {
            console.warn('Poll error for infra job request', job.reqId, err);
          }
        }

        for (const job of activeNlbJobs) {
          if (!job.nlbReqId) continue;
          try {
            let res = await beetleApi.getRequestDetails(job.nlbReqId, { brief: true });
            const isTerminalSuccess =
              res?.status === 'Success' || res?.status === 'Completed' || res?.status === 'Succeeded';
            if (isTerminalSuccess && res.responseData === undefined) {
              res = await beetleApi.getRequestDetails(job.nlbReqId, { brief: false });
            }
            if (res && res.status) {
              nlbStatusUpdates[job.id] = res;
            }
          } catch (err) {
            console.warn('Poll error for NLB job request', job.nlbReqId, err);
          }
        }

        setJobs((prevJobs) =>
          prevJobs.map((job) => {
            const isInfraDone = job.status === 'Success' || job.status === 'Failed';
            const isNlbDone = !job.nlbStatus || job.nlbStatus === 'Success' || job.nlbStatus === 'Failed';
            if (isInfraDone && isNlbDone) return job;

            const newElapsed = calculateJobElapsedSeconds(job);

            if (!job.isSample) {
              let updatedJob = { ...job, elapsedSeconds: newElapsed };

              if (job.nlbStatus === 'Provisioning' && job.nlbReqId) {
                const nlbUpdate = nlbStatusUpdates[job.id];
                if (nlbUpdate) {
                  if (nlbUpdate.status === 'Success' || nlbUpdate.status === 'Completed' || nlbUpdate.status === 'Succeeded') {
                    const createdNlb = nlbUpdate.responseData?.nlbList?.[0];
                    const nlbName =
                      createdNlb?.name ||
                      createdNlb?.id ||
                      job.targetNlbList?.[0]?.targetGroup?.nodeGroupId ||
                      job.targetNlbList?.[0]?.name ||
                      `${job.infraId}-nlb`;
                    updatedJob = {
                      ...updatedJob,
                      nlbStatus: 'Success',
                      nlbHealth: {
                        healthy: true,
                        status: 'Healthy',
                        checkedAt: new Date().toLocaleTimeString()
                      },
                      logs: [
                        ...updatedJob.logs,
                        `Stage 2 -> NLB Provisioned Successfully! (${nlbName})`,
                        `GET /beetle/request/${job.nlbReqId} -> Status: Success`
                      ]
                    };
                  } else if (nlbUpdate.status === 'Error' || nlbUpdate.status === 'Failed') {
                    updatedJob = {
                      ...updatedJob,
                      nlbStatus: 'Failed',
                      logs: [
                        ...updatedJob.logs,
                        `Stage 2 -> NLB Provisioning Failed: ${nlbUpdate.errorResponse || 'Backend NLB error'}`
                      ]
                    };
                  }
                }
              }

              if (job.status === 'Handling') {
                const update = statusUpdates[job.id];
                const realStatus = update?.status || 'Handling';
                if (update?.tumblebugSource) {
                  updatedJob.tumblebugSource = update.tumblebugSource;
                }
                if (update?.subRequestSummary) {
                  updatedJob.subRequestSummary = update.subRequestSummary;
                }
                if (update?.subRequests) {
                  updatedJob.subRequests = update.subRequests;
                }
                const liveNodes = liveNodeUpdates[job.id];
                if (Array.isArray(liveNodes) && liveNodes.length > 0) {
                  updatedJob.vms = liveNodes.map((n: any) => ({
                    name: n.name || n.id || 'node',
                    specId: n.specId || n.spec || 'c5.large',
                    publicIp: n.publicIP || n.publicIp || 'Allocating...',
                    privateIp: n.privateIP || n.privateIp || 'Allocating...',
                    status: n.status || 'Creating',
                    systemMessage: n.systemMessage || ''
                  }));
                }

                if (realStatus === 'Error' || realStatus === 'Failed') {
                  delete lastInfraNodePollAtRef.current[job.id];
                  const errorMsg = update?.errorResponse || 'Backend provisioning error';
                  const cleanedLogs = updatedJob.logs.filter((l) => !l.includes(`GET /beetle/request/${job.reqId}`));
                  return {
                    ...updatedJob,
                    status: 'Failed',
                    error: errorMsg,
                    logs: [
                      ...cleanedLogs,
                      `GET /beetle/request/${job.reqId}?brief=true -> Status: Error (${errorMsg}) (Duration: ${newElapsed}s)`
                    ]
                  };
                }

                if (realStatus === 'Success' || realStatus === 'Completed' || realStatus === 'Succeeded') {
                  delete lastInfraNodePollAtRef.current[job.id];
                  const shouldTriggerNlb =
                    job.targetNlbList && job.targetNlbList.length > 0 && (!job.nlbStatus || job.nlbStatus === 'Idle');
                  if (shouldTriggerNlb) {
                    beetleApi
                      .migrateNlb(job.nsId, job.infraId, job.targetNlbList!)
                      .then((nlbRes) => {
                        if (nlbRes.success && nlbRes.reqId) {
                          setJobs((prev) =>
                            prev.map((j) =>
                              j.id === job.id
                                ? {
                                    ...j,
                                    nlbStatus: 'Provisioning',
                                    nlbReqId: nlbRes.reqId,
                                    logs: [
                                      ...j.logs,
                                      `Stage 2 -> POST /beetle/migration/middleware/ns/${job.nsId}/infra/${job.infraId}/nlb`,
                                      `HTTP 202 Accepted (NLB ReqID: ${nlbRes.reqId}, Status: Provisioning)`
                                    ]
                                  }
                                : j
                            )
                          );
                        } else if (!nlbRes.success) {
                          setJobs((prev) =>
                            prev.map((j) =>
                              j.id === job.id
                                ? {
                                    ...j,
                                    nlbStatus: 'Failed',
                                    logs: [
                                      ...j.logs,
                                      `Stage 2 NLB Migration Failed: ${nlbRes.error || 'Unknown error'}`
                                    ]
                                  }
                                : j
                            )
                          );
                        }
                      })
                      .catch((err) => {
                        console.warn('NLB migration trigger error:', err);
                      });
                  }

                  const rawNodes = update?.responseData?.node || update?.responseData?.infraInfo?.node || [];
                  const parsedVms =
                    Array.isArray(rawNodes) && rawNodes.length > 0
                      ? rawNodes.map((n: any) => ({
                          name: n.name || n.id || 'node',
                          specId: n.specId || n.spec || 'c5.large',
                          publicIp: n.publicIP || n.publicIp || 'N/A',
                          privateIp: n.privateIP || n.privateIp || 'N/A',
                          status:
                            n.status || ((n.publicIP || n.publicIp || n.privateIP || n.privateIp) ? 'Running' : 'Failed'),
                          systemMessage: n.systemMessage || ''
                        }))
                      : updatedJob.vms;

                  const cleanedLogs = updatedJob.logs.filter((l) => !l.includes(`GET /beetle/request/${job.reqId}`));
                  return {
                    ...updatedJob,
                    status: 'Success',
                    nlbStatus: shouldTriggerNlb ? 'Provisioning' : updatedJob.nlbStatus,
                    vms: parsedVms,
                    logs: [
                      ...cleanedLogs,
                      `GET /beetle/request/${job.reqId} -> Status: Success (Duration: ${newElapsed}s)`
                    ]
                  };
                }

                const stepProgress = updatedJob.subRequestSummary
                  ? ` [${updatedJob.subRequestSummary.succeeded}/${updatedJob.subRequestSummary.total} steps${updatedJob.subRequestSummary.currentAction ? `: ${updatedJob.subRequestSummary.currentAction}` : ''}]`
                  : '';
                return {
                  ...updatedJob,
                  logs: [
                    ...updatedJob.logs.filter((l) => !l.includes(`GET /beetle/request/${job.reqId}`)),
                    `GET /beetle/request/${job.reqId}?brief=true -> Status: Handling${stepProgress} (Elapsed: ${newElapsed}s)`
                  ]
                };
              }

              return updatedJob;
            }
            return job;
          })
        );
      }

      // 2. Poll Active Object Storage Jobs
      if (activeOsJobs.length > 0) {
        for (const job of activeOsJobs) {
          if (!job.reqId || job.isSample) continue;
          try {
            const reqDetails = await beetleApi.getRequestDetails(job.reqId, { brief: true });
            const status = reqDetails?.status;
            if (status === 'Completed' || status === 'Succeeded' || status === 'Success') {
              setObjectStorageJobs((prevJobs) =>
                prevJobs.map((j) => {
                  if (j.id === job.id) {
                    const dur = calculateJobElapsedSeconds(j);
                    return {
                      ...j,
                      status: 'Success',
                      elapsedSeconds: dur,
                      logs: [
                        ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                        `GET /beetle/request/${job.reqId} -> Status: Success (Duration: ${dur}s)`
                      ]
                    };
                  }
                  return j;
                })
              );
            } else if (status === 'Failed' || status === 'Error') {
              setObjectStorageJobs((prevJobs) =>
                prevJobs.map((j) => {
                  if (j.id === job.id) {
                    const dur = calculateJobElapsedSeconds(j);
                    return {
                      ...j,
                      status: 'Failed',
                      elapsedSeconds: dur,
                      logs: [
                        ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                        `GET /beetle/request/${job.reqId} -> Status: Failed (${reqDetails?.errorResponse || 'Execution Error'}) (Duration: ${dur}s)`
                      ]
                    };
                  }
                  return j;
                })
              );
            } else {
              setObjectStorageJobs((prevJobs) =>
                prevJobs.map((j) => (j.id === job.id ? { ...j, elapsedSeconds: calculateJobElapsedSeconds(j) } : j))
              );
            }
          } catch (err) {
            console.warn('Poll error for OS job request', job.reqId, err);
          }
        }
      }

      // 3. Poll Active Data Migration Jobs (Every 5 seconds per job using lightweight brief=true)
      if (activeDataJobs.length > 0) {
        const now = Date.now();
        for (const job of activeDataJobs) {
          if (!job.reqId || job.isSample) continue;
          const lastPolled = lastDataPollAtRef.current[job.id] ?? job.createdAtMs ?? 0;
          if (now - lastPolled < 5000) {
            continue;
          }
          lastDataPollAtRef.current[job.id] = now;
          try {
            let reqDetails = await beetleApi.getRequestDetails(job.reqId, { brief: true });
            const status = reqDetails?.status;
            if (status === 'Completed' || status === 'Succeeded' || status === 'Success') {
              delete lastDataPollAtRef.current[job.id];
              if (reqDetails.responseData === undefined) {
                reqDetails = await beetleApi.getRequestDetails(job.reqId, { brief: false });
              }
              const resData = reqDetails?.responseData || {};
              const migResult = resData?.result || resData;
              setDataJobs((prevJobs) =>
                prevJobs.map((j) => {
                  if (j.id === job.id) {
                    const dur = calculateJobElapsedSeconds(j);
                    const detailMsg = resData?.message ? ` - ${resData.message}` : '';
                    return {
                      ...j,
                      status: 'Success',
                      elapsedSeconds: dur,
                      hookOutputs: migResult?.hookOutputs || j.hookOutputs,
                      bytesTransferred: typeof migResult?.bytesTransferred === 'number' ? migResult.bytesTransferred : j.bytesTransferred,
                      filesTransferred: typeof migResult?.filesTransferred === 'number' ? migResult.filesTransferred : j.filesTransferred,
                      logs: [
                        ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                        `GET /beetle/request/${job.reqId} -> Status: Success${detailMsg} (Duration: ${dur}s)`
                      ]
                    };
                  }
                  return j;
                })
              );
            } else if (status === 'Failed' || status === 'Error') {
              delete lastDataPollAtRef.current[job.id];
              setDataJobs((prevJobs) =>
                prevJobs.map((j) => {
                  if (j.id === job.id) {
                    const dur = calculateJobElapsedSeconds(j);
                    return {
                      ...j,
                      status: 'Failed',
                      elapsedSeconds: dur,
                      logs: [
                        ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                        `GET /beetle/request/${job.reqId} -> Status: Failed (${reqDetails?.errorResponse || 'Execution Error'}) (Duration: ${dur}s)`
                      ]
                    };
                  }
                  return j;
                })
              );
            } else {
              setDataJobs((prevJobs) =>
                prevJobs.map((j) => {
                  if (j.id === job.id) {
                    const dur = calculateJobElapsedSeconds(j);
                    return {
                      ...j,
                      elapsedSeconds: dur,
                      logs: [
                        ...j.logs.filter((l) => !l.includes('GET /beetle/request/')),
                        `GET /beetle/request/${job.reqId}?brief=true -> Status: Handling (Elapsed: ${dur}s)`
                      ]
                    };
                  }
                  return j;
                })
              );
            }
          } catch (err) {
            console.warn('Poll error for data job request', job.reqId, err);
          }
        }
      }
    }, 2500);

    return () => clearInterval(interval);
  }, [setJobs, setObjectStorageJobs, setDataJobs]);
}
