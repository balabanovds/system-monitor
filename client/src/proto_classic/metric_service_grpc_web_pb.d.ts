import * as grpcWeb from 'grpc-web';

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

import {
  Metric,
  ParsersInfoResponse,
  Request} from './metric_service_pb';

export class MetricsClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  getStream(
    request: Request,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Metric>;

  parsersInfo(
    request: google_protobuf_empty_pb.Empty,
    metadata: grpcWeb.Metadata | undefined,
    callback: (err: grpcWeb.Error,
               response: ParsersInfoResponse) => void
  ): grpcWeb.ClientReadableStream<ParsersInfoResponse>;

}

export class MetricsPromiseClient {
  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: string; });

  getStream(
    request: Request,
    metadata?: grpcWeb.Metadata
  ): grpcWeb.ClientReadableStream<Metric>;

  parsersInfo(
    request: google_protobuf_empty_pb.Empty,
    metadata?: grpcWeb.Metadata
  ): Promise<ParsersInfoResponse>;

}

