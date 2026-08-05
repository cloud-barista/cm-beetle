'use client';

import { useEffect, useRef } from 'react';
import { useMigrationStore, calculateJobElapsedSeconds } from '../store/migrationStore';
import { beetleApi } from '../api/client';

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
      const activeOsJobs = currentOsJobs.filter((j) => !j.isSample && j.status === 'Handling');
      const activeDataJobs = currentDataJobs.filter((j) => !j.isSample && j.status === 'Handling');

      if (activeInfraJobs.length === 0 && activeOsJobs.length === 0 && activeDataJobs.length === 0) {
        return;
      }

      // 1. Poll Active Infrastructure Migration Jobs
      if (activeInfraJobs.length > 0) {
        const statusUpdates: Record<string, { status: string; errorResponse?: string; responseData?: any }> = {};

        for (const job of activeInfraJobs) {
          try {
            const res = await beetleApi.getRequestDetails(job.reqId);
            if (res && res.status) {
              statusUpdates[job.id] = {
                status: res.status,
                errorResponse: res.errorResponse,
                responseData: res.responseData
              };
            }
          } catch (err) {
            console.warn('Poll error for infra job request', job.reqId, err);
          }
        }

        setJobs((prevJobs) =>
          prevJobs.map((job) => {
            if (job.status === 'Success' || job.status === 'Failed') return job;

            const newElapsed = calculateJobElapsedSeconds(job);

            if (!job.isSample) {
              const update = statusUpdates[job.id];
              const realStatus = update?.status || 'Handling';

              if (realStatus === 'Error' || realStatus === 'Failed') {
                const errorMsg = update?.errorResponse || 'Backend provisioning error';
                const cleanedLogs = job.logs.filter((l) => !l.includes('GET /beetle/request/'));
                return {
                  ...job,
                  status: 'Failed',
                  elapsedSeconds: newElapsed,
                  error: errorMsg,
                  logs: [
                    ...cleanedLogs,
                    `GET /beetle/request/${job.reqId} -> Status: Error (${errorMsg}) (Duration: ${newElapsed}s)`
                  ]
                };
              }

              if (realStatus === 'Success' || realStatus === 'Completed' || realStatus === 'Succeeded') {
                const rawNodes = update?.responseData?.node || update?.responseData?.infraInfo?.node || [];
                const parsedVms = Array.isArray(rawNodes) && rawNodes.length > 0
                  ? rawNodes.map((n: any) => ({
                      name: n.name || n.id || 'node',
                      specId: n.specId || n.spec || 'c5.large',
                      publicIp: n.publicIP || n.publicIp || 'N/A',
                      privateIp: n.privateIP || n.privateIp || 'N/A'
                    }))
                  : job.vms;

                const cleanedLogs = job.logs.filter((l) => !l.includes('GET /beetle/request/'));
                return {
                  ...job,
                  status: 'Success',
                  elapsedSeconds: newElapsed,
                  vms: parsedVms,
                  logs: [
                    ...cleanedLogs,
                    `GET /beetle/request/${job.reqId} -> Status: Success (Duration: ${newElapsed}s)`
                  ]
                };
              }

              return {
                ...job,
                elapsedSeconds: newElapsed,
                logs: [
                  ...job.logs.filter((l) => !l.includes('GET /beetle/request/')),
                  `GET /beetle/request/${job.reqId} -> Status: Handling (Elapsed: ${newElapsed}s)`
                ]
              };
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
            const reqDetails = await beetleApi.getRequestDetails(job.reqId);
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

      // 3. Poll Active Data Migration Jobs
      if (activeDataJobs.length > 0) {
        for (const job of activeDataJobs) {
          if (!job.reqId || job.isSample) continue;
          try {
            const reqDetails = await beetleApi.getRequestDetails(job.reqId);
            const status = reqDetails?.status;
            if (status === 'Completed' || status === 'Succeeded' || status === 'Success') {
              setDataJobs((prevJobs) =>
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
                prevJobs.map((j) => (j.id === job.id ? { ...j, elapsedSeconds: calculateJobElapsedSeconds(j) } : j))
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
