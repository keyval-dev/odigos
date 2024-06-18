import { getNodeAutoInstrumentations } from "@opentelemetry/auto-instrumentations-node";
import { OTLPTraceExporter } from "@opentelemetry/exporter-trace-otlp-grpc";
import { NodeSDK } from "@opentelemetry/sdk-node";
import { OpAMPClientHttp } from "./opamp";
import {
  SEMRESATTRS_TELEMETRY_SDK_LANGUAGE,
  TELEMETRYSDKLANGUAGEVALUES_NODEJS,
  SEMRESATTRS_TELEMETRY_SDK_NAME,
  SEMRESATTRS_TELEMETRY_SDK_VERSION,
  SEMRESATTRS_PROCESS_PID,
} from "@opentelemetry/semantic-conventions";
import { diag } from "@opentelemetry/api";

const opampServerHost = process.env.ODIGOS_OPAMP_SERVER_HOST;
const instrumentationDeviceId = process.env.ODIGOS_INSTRUMENTATION_DEVICE_ID;
if (!opampServerHost || !instrumentationDeviceId) {
  throw new Error(
    "ODIGOS_OPAMP_SERVER_HOST and ODIGOS_INSTRUMENTATION_DEVICE_ID must be set"
  );
}

if (opampServerHost) {
  const opampClient = new OpAMPClientHttp({
    instrumentationDeviceId: instrumentationDeviceId,
    opAMPServerHost: opampServerHost,
    agentDescriptionIdentifyingAttributes: {
      [SEMRESATTRS_TELEMETRY_SDK_LANGUAGE]: TELEMETRYSDKLANGUAGEVALUES_NODEJS,
      [SEMRESATTRS_TELEMETRY_SDK_NAME]: "odigos",
      [SEMRESATTRS_TELEMETRY_SDK_VERSION]: "0.0.1", // TODO: get version from package.json
      [SEMRESATTRS_PROCESS_PID]: process.pid,
    },
    agentDescriptionNonIdentifyingAttributes: {},
  });

  opampClient.start();

  const sdk = new NodeSDK({
    resourceDetectors: [opampClient],
    instrumentations: [getNodeAutoInstrumentations()],
    traceExporter: new OTLPTraceExporter(),
  });
  sdk.start();

  const shutdown = async () => {
    try {
      diag.debug('Shutting down OpenTelemetry SDK and OpAMP client');
      await Promise.all([sdk.shutdown(), opampClient.shutdown()]);
      await new Promise(resolve => setTimeout(resolve, 60000));
    } catch(err) {
      diag.error('Error shutting down OpenTelemetry SDK and OpAMP client', err);
    }
  }

  process.on('SIGTERM', shutdown);  
  process.on('SIGINT', shutdown);
  process.on('exit', shutdown);

} else {
  const sdk = new NodeSDK({
    autoDetectResources: true,
    instrumentations: [getNodeAutoInstrumentations()],
    traceExporter: new OTLPTraceExporter(),
  });  
  sdk.start();
}


