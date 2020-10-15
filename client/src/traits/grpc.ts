import { grpc } from "@improbable-eng/grpc-web";
import { Metric, Request } from "../proto/metric_service_pb";
import { Metrics } from "../proto/metric_service_pb_service";

import { ref } from "vue";
import { HOST, PORT } from "@/main";

export interface GMetric {
  title: string;
  type: number;
  seconds: number;
  nanos: number;
  value: number;
}

export class GrpcMetricClient {
  private n = ref(0);
  private m = ref(0);
  private data = ref(Array<GMetric>(0));
  private active = ref(false);
  private addr: string;

  private streamClient?: grpc.Client<Request, Metric> = undefined;

  private static instance: GrpcMetricClient;

  private constructor(host: string, port: number) {
    this.addr = `http://${host}:${port}`;
  }

  public static getInstance(): GrpcMetricClient {
    if (!GrpcMetricClient.instance) {
      GrpcMetricClient.instance = new GrpcMetricClient(HOST, PORT);
    }
    return GrpcMetricClient.instance;
  }

  public startStream() {
    if (this.n.value === 0 || this.m.value === 0) {
      throw new Error("n and m values should be positive");
    }
    const req = new Request();
    req.setN(this.n.value);
    req.setM(this.m.value);

    this.streamClient = grpc.client(Metrics.GetStream, {
      host: this.addr,
    });

    // this.streamClient.onHeaders((headers) => {
    //   console.log("headers", headers);
    // });

    this.streamClient.onMessage((msg) => {
      const m = msg.toObject();
      const t = m.time;
      if (!t) {
        return;
      }
      const gMetric: GMetric = {
        title: m.title,
        type: m.type,
        seconds: t.seconds,
        nanos: t.nanos,
        value: m.value,
      };
      this.data.value.push(gMetric);
    });
    this.streamClient.onEnd((code: grpc.Code, msg: string) => {
      console.log(`closed stream. code ${code}, message: ${msg}`, code, msg);
      this.streamClient?.close();
      this.active.value = false;
    });
    this.streamClient.start();

    this.active.value = true;
    this.streamClient.send(req);
  }

  public stopStream() {
    this.streamClient?.close();
    this.active.value = false;
  }

  public getters() {
    return {
      n: this.n,
      m: this.m,
      data: this.data,
      active: this.active,
    };
  }
}

export const getStream = (host: string, port: number) => {
  const n = ref(0);
  const m = ref(0);

  const data = ref(Array<Metric.AsObject>(0));

  const addr = `http://${host}:${port}`;

  let client: grpc.Client<Request, Metric>;

  function start() {
    const req = new Request();
    req.setN(n.value);
    req.setM(m.value);

    client = grpc.client(Metrics.GetStream, {
      host: addr,
    });

    // client.onHeaders((headers: grpc.Metadata) => {
    //     console.log('Metrics.GetStream.onHeaders', headers);
    // });
    client.onMessage((msg) => {
      console.log("new metric", msg.toObject());
      data.value.push(msg.toObject());
    });
    // client.onEnd(
    //     (code: grpc.Code, msg: string, trailers: grpc.Metadata) => {
    //         console.log('Metrics.GetStream.onEnd', code, msg, trailers);
    //     },
    // );
    client.start();
    client.send(req);
  }

  function stop() {
    client.close();
  }

  return {
    n,
    m,
    data,
    start,
    stop,
  };
};

type Done = {
  code: string;
  message: string;
};
