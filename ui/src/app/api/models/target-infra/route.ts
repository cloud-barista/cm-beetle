import { NextResponse } from 'next/server';
import fs from 'fs';
import path from 'path';

const standardFilePath = path.join(process.cwd(), 'src', 'data', 'sampleTargetInfra.json');
const gpuFilePath = path.join(process.cwd(), 'src', 'data', 'sampleTargetGpuInfra.json');

const standardPublicPath = path.join(process.cwd(), 'public', 'models', 'target_infra_model.json');
const gpuPublicPath = path.join(process.cwd(), 'public', 'models', 'target_gpu_infra_model.json');

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
    const isGpu = searchParams.get('sample') === 'gpu' || body.isGpuSample || body.modelId?.includes('gpu');
    const targetFile = isGpu ? gpuFilePath : standardFilePath;
    const publicPath = isGpu ? gpuPublicPath : standardPublicPath;

    let existingData: any = { success: true, data: [] };

    if (fs.existsSync(targetFile)) {
      try {
        const fileContent = fs.readFileSync(targetFile, 'utf8');
        existingData = JSON.parse(fileContent);
      } catch {
        existingData = { success: true, data: [] };
      }
    }

    let finalData: any;
    if (body.data && Array.isArray(body.data)) {
      finalData = {
        ...existingData,
        ...body,
      };
    } else if (Array.isArray(body)) {
      finalData = {
        success: true,
        data: body,
      };
    } else if (body.targetInfra || body.targetVNet) {
      // Single recommendation candidate edited
      if (Array.isArray(existingData.data) && existingData.data.length > 0) {
        existingData.data[0] = {
          ...existingData.data[0],
          ...body,
        };
        finalData = existingData;
      } else {
        finalData = {
          success: true,
          data: [body],
        };
      }
    } else {
      finalData = {
        ...existingData,
        ...body,
      };
    }

    // Clean up temporary flags
    delete finalData.isGpuSample;
    delete finalData.modelId;

    // Ensure dir exists
    const dir = path.dirname(targetFile);
    if (!fs.existsSync(dir)) {
      fs.mkdirSync(dir, { recursive: true });
    }

    fs.writeFileSync(targetFile, JSON.stringify(finalData, null, 2), 'utf8');

    // Also sync to public/models as a CloudModelEnvelope
    try {
      const pubDir = path.dirname(publicPath);
      if (!fs.existsSync(pubDir)) fs.mkdirSync(pubDir, { recursive: true });

      const candidate = finalData.data?.[0] || finalData;
      const modelEnvelope = {
        id: isGpu ? 'model-target-gpu-infra-1' : 'model-target-infra-1',
        name: isGpu ? 'Recommended AWS AI GPU Cluster' : 'Recommended AWS Web & InfluxDB Cloud Infra',
        description: isGpu
          ? 'Optimized AWS cloud GPU infrastructure model (ap-northeast-2) recommended for heterogeneous GPU cluster'
          : 'Optimized AWS cloud infrastructure model (ap-northeast-2) recommended for standard on-premise cluster',
        modelType: 'cloud',
        version: '1.0',
        cloudInfraModel: candidate,
        updatedTime: new Date().toISOString(),
      };
      fs.writeFileSync(publicPath, JSON.stringify(modelEnvelope, null, 2), 'utf8');
    } catch {
      // Non-critical
    }

    const fileName = path.basename(targetFile);
    return NextResponse.json({
      success: true,
      message: `Target cloud infrastructure model saved and overwritten to ${fileName} successfully.`,
      fileName,
    });
  } catch (error: any) {
    return NextResponse.json({ success: false, error: error.message }, { status: 500 });
  }
}
