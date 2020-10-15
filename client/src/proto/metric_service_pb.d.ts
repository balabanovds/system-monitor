// package: metric
// file: metric_service.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";

export class Metric extends jspb.Message {
  hasTime(): boolean;
  clearTime(): void;
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): void;

  getType(): MetricTypeMap[keyof MetricTypeMap];
  setType(value: MetricTypeMap[keyof MetricTypeMap]): void;

  getTitle(): string;
  setTitle(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metric.AsObject;
  static toObject(includeInstance: boolean, msg: Metric): Metric.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Metric, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metric;
  static deserializeBinaryFromReader(message: Metric, reader: jspb.BinaryReader): Metric;
}

export namespace Metric {
  export type AsObject = {
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    type: MetricTypeMap[keyof MetricTypeMap],
    title: string,
    value: number,
  }
}

export class Request extends jspb.Message {
  getN(): number;
  setN(value: number): void;

  getM(): number;
  setM(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Request.AsObject;
  static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Request, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Request;
  static deserializeBinaryFromReader(message: Request, reader: jspb.BinaryReader): Request;
}

export namespace Request {
  export type AsObject = {
    n: number,
    m: number,
  }
}

export interface MetricTypeMap {
  UNDEFINED: 0;
  LOADAVERAGE1MIN: 1;
  LOADAVERAGE5MIN: 2;
  LOADAVERAGE15MIN: 3;
  CPUUSER: 4;
  CPUSYSTEM: 5;
  CPUIDLE: 6;
  IOTPS: 7;
  IOREADKBPS: 8;
  IOWRITEKBPS: 9;
  IOCPUUSER: 10;
  IOCPUSYSTEM: 11;
  IOCPUIDLE: 12;
  FSMBFREE: 13;
  FSINODEFREE: 14;
}

export const MetricType: MetricTypeMap;

