import {grpc} from '@improbable-eng/grpc-web';
import {Metric, ParsersInfoResponse, Request,} from '@/proto/metric_service_pb';
import {Metrics} from '@/proto/metric_service_pb_service';
import {Empty} from 'google-protobuf/google/protobuf/empty_pb';

import {ref} from 'vue';
import {HOST, PORT} from '@/consts';

export interface GMetric {
    title: string;
    type: number;
    seconds: number;
    nanos: number;
    value: number;
}

export interface ChartInfo {
    title: string;
    metricTypes: number[];
}

export class GrpcMetricClient {
    private n = ref(0);
    private m = ref(0);
    private data = ref(Array<GMetric>(0));
    private infoList = ref(Array<ChartInfo>(0));

    private active = ref(false);
    private error = ref('');
    private readonly addr: string;

    private streamClient?: grpc.Client<Request, Metric> = undefined;
    private infoClient?: grpc.Client<Empty, ParsersInfoResponse> = undefined;

    private static instance: GrpcMetricClient;

    private constructor(host: string, port: number) {
        this.addr = `http://${host}:${port}`;
    }

    private async ping() {
        await fetch(`http://${HOST}:${PORT}/metric.Metrics/ParsersInfo`, {
            method: 'POST',
            mode: 'no-cors'
        })
            .catch(e => {
                this.error.value = 'server is not reachable'
            });
    }

    public static getInstance(): GrpcMetricClient {
        if (!GrpcMetricClient.instance) {
            GrpcMetricClient.instance = new GrpcMetricClient(HOST, PORT);
        }
        return GrpcMetricClient.instance;
    }

    public startStream() {
        this.error.value = ''
        if (this.n.value === 0 || this.m.value === 0) {
            this.error.value = 'values should be >= 0';
            this.active.value = false;
            return
        }
        const req = new Request();
        req.setN(this.n.value);
        req.setM(this.m.value);

        this.streamClient = grpc.client(Metrics.GetStream, {
            host: this.addr,
        });

        this.streamClient.onMessage(msg => {
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
            console.log(
                `closed stream. code ${code}, message: ${msg}`,
                code,
                msg,
            );
            this.error.value = "GRPC server closed connection"
            this.streamClient?.close();
            this.active.value = false;
        });
        this.streamClient.start();

        this.active.value = true;
        this.streamClient.send(req);
    }

    public stopStream() {
        try {
            this.streamClient?.close();
        } catch (e) {
            this.error.value = 'failed to stop grpc client'
            console.error('failed to stop grpc client', e);
        } finally {
            this.active.value = false;
        }
    }

    public getInfo() {
        this.ping()
        if (this.error.value) {
            return
        }
        this.infoClient = grpc.client(Metrics.ParsersInfo, {
            host: this.addr,
        });

        this.infoClient.onMessage((msg: ParsersInfoResponse) => {
            for (const p of msg.getListList()) {
                const gInfo: ChartInfo = {
                    title: p.getName(),
                    metricTypes: p.getMetrictypesList(),
                };
                this.infoList.value.push(gInfo);
            }
        });

        this.infoClient.start();
        this.infoClient.send(new Empty());
    }

    public closeInfoClient() {
        this.infoClient?.close();
    }

    public streamGetters() {
        return {
            n: this.n,
            m: this.m,
            data: this.data,
            active: this.active,
        };
    }

    public infoGetter() {
        return {
            list: this.infoList,
            error: this.error
        };
    }
}
