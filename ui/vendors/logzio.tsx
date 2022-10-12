import {
  ObservabilityVendor,
  ObservabilitySignals,
  VendorObjects,
  VendorType,
} from "@/vendors/index";
import Logziologo from "@/img/vendor/logzio.svg";
import { NextApiRequest } from "next";

export class Logzio implements ObservabilityVendor {
  name = "logzio";
  displayName = "Logz.io";
  type = VendorType.MANAGED;
  supportedSignals = [
    ObservabilitySignals.Metrics,
    ObservabilitySignals.Traces,
    ObservabilitySignals.Logs,
  ];

  getLogo = (props: any) => {
    return <Logziologo {...props} />;
  };

  getFields = (selectedSignals: any) => {
    let fields = [
      {
        displayName: "Region",
        id: "region",
        name: "region",
        type: "text",
      },
    ];
    if (selectedSignals[ObservabilitySignals.Traces]) {
      fields.push({
        displayName: "Tracing token",
        id: "tracingToken",
        name: "tracingToken",
        type: "password",
      });
    }
    if (selectedSignals[ObservabilitySignals.Metrics]) {
      fields.push({
        displayName: "Metrics token",
        id: "metricsToken",
        name: "metricsToken",
        type: "password",
      });
    }
    if (selectedSignals[ObservabilitySignals.Logs]) {
      fields.push({
        displayName: "Logs token",
        id: "logsToken",
        name: "logsToken",
        type: "password",
      });
    }
    return fields;
  };

  toObjects = (req: NextApiRequest) => {
    let metricsToken, logsToken, tracingToken;
    if (req.body.metricsToken) {
      metricsToken = Buffer.from(req.body.metricsToken).toString("base64");
    } else {
      metricsToken = "none";
    }
    if (req.body.logsToken) {
      logsToken = Buffer.from(req.body.logsToken).toString("base64");
    } else {
      logsToken = "none";
    }
    if (req.body.tracingToken) {
      tracingToken = Buffer.from(req.body.tracingToken).toString("base64");
    } else {
      tracingToken = "none";
    }
    return {
      Data: {
        LOGZIO_REGION: req.body.region,
      },
      Secret: {
        LOGZIO_TRACING_TOKEN: tracingToken,
        LOGZIO_LOGS_TOKEN: logsToken,
        LOGZIO_METRICS_TOKEN: metricsToken,
      },
    };
  };
  mapDataToFields = (data: any) => {
    return {
      region: data.LOGZIO_REGION || "us",
    };
  };
}
