// package: metric
// file: metric_service.proto

import * as jspb from "google-protobuf";
import * as google_protobuf_timestamp_pb from "google-protobuf/google/protobuf/timestamp_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";

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

export class ParsersInfoResponse extends jspb.Message {
  clearListList(): void;
  getListList(): Array<ParserInfo>;
  setListList(value: Array<ParserInfo>): void;
  addList(value?: ParserInfo, index?: number): ParserInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParsersInfoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ParsersInfoResponse): ParsersInfoResponse.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ParsersInfoResponse, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParsersInfoResponse;
  static deserializeBinaryFromReader(message: ParsersInfoResponse, reader: jspb.BinaryReader): ParsersInfoResponse;
}

export namespace ParsersInfoResponse {
  export type AsObject = {
    listList: Array<ParserInfo.AsObject>,
  }
}

export class ParserInfo extends jspb.Message {
  getType(): ParserTypeMap[keyof ParserTypeMap];
  setType(value: ParserTypeMap[keyof ParserTypeMap]): void;

  clearMetrictypesList(): void;
  getMetrictypesList(): Array<MetricTypeMap[keyof MetricTypeMap]>;
  setMetrictypesList(value: Array<MetricTypeMap[keyof MetricTypeMap]>): void;
  addMetrictypes(value: MetricTypeMap[keyof MetricTypeMap], index?: number): MetricTypeMap[keyof MetricTypeMap];

  getName(): string;
  setName(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParserInfo.AsObject;
  static toObject(includeInstance: boolean, msg: ParserInfo): ParserInfo.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ParserInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParserInfo;
  static deserializeBinaryFromReader(message: ParserInfo, reader: jspb.BinaryReader): ParserInfo;
}

export namespace ParserInfo {
  export type AsObject = {
    type: ParserTypeMap[keyof ParserTypeMap],
    metrictypesList: Array<MetricTypeMap[keyof MetricTypeMap]>,
    name: string,
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

export interface ParserTypeMap {
  UNDEF: 0;
  LOADAVERAGE: 1;
  CPU: 2;
  IO: 3;
  FS: 4;
  NET: 5;
}

export const ParserType: ParserTypeMap;

