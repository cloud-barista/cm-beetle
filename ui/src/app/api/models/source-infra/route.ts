import { NextResponse } from 'next/server';
import fs from 'fs';
import path from 'path';

const standardFilePath = path.join(process.cwd(), 'src', 'data', 'sampleSourceInfra.json');
const gpuFilePath = path.join(process.cwd(), 'src', 'data', 'sampleSourceGpuInfra.json');

const standardPublicPath = path.join(process.cwd(), 'public', 'models', 'source_infra_model.json');
const gpuPublicPath = path.join(process.cwd(), 'public', 'models', 'source_gpu_infra_model.json');

export async function GET(request: Request) {
  try {
    const { searchParams } = new URL(request.url);
    const isGpu = searchParams.get('sample') === 'gpu' || searchParams.get('type') === 'gpu';
    const targetFile = isGpu ? gpuFilePath : standardFilePath;

    if (!fs.existsSync(targetFile)) {
      return NextResponse.json({ success: false, error: 'File not found' }, { status: 404 });
    }
    const fileContent = fs.readFileSync(targetFile, 'utf8');
    const data = JSON.parse(fileContent);
    return NextResponse.json({ success: true, data, sampleType: isGpu ? 'gpu' : 'standard' });
  } catch (error: any) {
    return NextResponse.json({ success: false, error: error.message }, { status: 500 });
  }
}

export async function POST(request: Request) {
  try {
    const { searchParams } = new URL(request.url);
    const body = await request.json();
    const isGpu = searchParams.get('sample') === 'gpu' || body.isGpuSample || body.modelId === 'sample-source-gpu-infra-1';
    const targetFile = isGpu ? gpuFilePath : standardFilePath;
    const publicPath = isGpu ? gpuPublicPath : standardPublicPath;

    let existingData: any = {};
    if (fs.existsSync(targetFile)) {
      try {
        const fileContent = fs.readFileSync(targetFile, 'utf8');
        existingData = JSON.parse(fileContent);
      } catch {
        existingData = {};
      }
    }

    let finalData: any;
    if (body.sourceInfra) {
      finalData = {
        ...existingData,
        ...body,
        sourceInfra: body.sourceInfra,
      };
    } else if (body.nodes || body.network) {
      // Direct OnpremInfra object
      finalData = {
        ...existingData,
        sourceInfra: body,
      };
    } else {
      finalData = {
        ...existingData,
        ...body,
      };
    }

    // Clean up temporary control flags before saving to JSON file
    delete finalData.isGpuSample;
    delete finalData.modelId;

    // Ensure dir exists
    const dir = path.dirname(targetFile);
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }

    fs.writeFileSync(targetFile, JSON.stringify(finalData, null, 2), 'utf8');

    // Also sync to public/models as an OnpremModelEnvelope
    try {
      const pubDir = path.dirname(publicPath);
      if (!fs.existsSync(pubDir)) fs.mkdirSync(pubDir, { recursive: true });

      const onpremiseInfraModel = finalData.sourceInfra || finalData;
      const modelEnvelope = {
        id: isGpu ? 'model-source-gpu-infra-1' : 'model-source-infra-1',
        name: isGpu ? 'AI Heterogeneous GPU Cluster Model' : 'Standard Web HAProxy & InfluxDB Cluster',
        description: isGpu
          ? 'On-premise AI training and inference cluster with heterogeneous GPUs (2x A100 80GB, 1x T4, 1x MI350, 1x L4)'
          : 'On-premise infrastructure model with 1 HAProxy load balancer and 2 InfluxDB database nodes (GPU-less)',
        modelType: 'onprem',
        version: '1.0',
        onpremiseInfraModel,
        updatedTime: new Date().toISOString(),
      };
      fs.writeFileSync(publicPath, JSON.stringify(modelEnvelope, null, 2), 'utf8');
    } catch {
      // Non-critical
    }

    const fileName = path.basename(targetFile);
    return NextResponse.json({
      success: true,
      message: `Source infrastructure model saved and overwritten to ${fileName} successfully.`,
      fileName,
    });
  } catch (error: any) {
    return NextResponse.json({ success: false, error: error.message }, { status: 500 });
  }
}
