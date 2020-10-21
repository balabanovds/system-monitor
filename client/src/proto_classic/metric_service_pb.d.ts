import * as jspb from "google-protobuf"

import * as google_protobuf_timestamp_pb from 'google-protobuf/google/protobuf/timestamp_pb';
import * as google_protobuf_empty_pb from 'google-protobuf/google/protobuf/empty_pb';

export class Metric extends jspb.Message {
  getTime(): google_protobuf_timestamp_pb.Timestamp | undefined;
  setTime(value?: google_protobuf_timestamp_pb.Timestamp): Metric;
  hasTime(): boolean;
  clearTime(): Metric;

  getType(): MetricType;
  setType(value: MetricType): Metric;

  getTitle(): string;
  setTitle(value: string): Metric;

  getValue(): number;
  setValue(value: number): Metric;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Metric.AsObject;
  static toObject(includeInstance: boolean, msg: Metric): Metric.AsObject;
  static serializeBinaryToWriter(message: Metric, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Metric;
  static deserializeBinaryFromReader(message: Metric, reader: jspb.BinaryReader): Metric;
}

export namespace Metric {
  export type AsObject = {
    time?: google_protobuf_timestamp_pb.Timestamp.AsObject,
    type: MetricType,
    title: string,
    value: number,
  }
}

export class Request extends jspb.Message {
  getN(): number;
  setN(value: number): Request;

  getM(): number;
  setM(value: number): Request;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Request.AsObject;
  static toObject(includeInstance: boolean, msg: Request): Request.AsObject;
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
  getListList(): Array<ParserInfo>;
  setListList(value: Array<ParserInfo>): ParsersInfoResponse;
  clearListList(): ParsersInfoResponse;
  addList(value?: ParserInfo, index?: number): ParserInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParsersInfoResponse.AsObject;
  static toObject(includeInstance: boolean, msg: ParsersInfoResponse): ParsersInfoResponse.AsObject;
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
  getType(): ParserType;
  setType(value: ParserType): ParserInfo;

  getMetrictypesList(): Array<MetricType>;
  setMetrictypesList(value: Array<MetricType>): ParserInfo;
  clearMetrictypesList(): ParserInfo;
  addMetrictypes(value: MetricType, index?: number): ParserInfo;

  getName(): string;
  setName(value: string): ParserInfo;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ParserInfo.AsObject;
  static toObject(includeInstance: boolean, msg: ParserInfo): ParserInfo.AsObject;
  static serializeBinaryToWriter(message: ParserInfo, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ParserInfo;
  static deserializeBinaryFromReader(message: ParserInfo, reader: jspb.BinaryReader): ParserInfo;
}

export namespace ParserInfo {
  export type AsObject = {
    type: ParserType,
    metrictypesList: Array<MetricType>,
    name: string,
  }
}

export enum MetricType { 
  UNDEFINED = 0,
  LOADAVERAGE1MIN = 1,
  LOADAVERAGE5MIN = 2,
  LOADAVERAGE15MIN = 3,
  CPUUSER = 4,
  CPUSYSTEM = 5,
  CPUIDLE = 6,
  IOTPS = 7,
  IOREADKBPS = 8,
  IOWRITEKBPS = 9,
  IOCPUUSER = 10,
  IOCPUSYSTEM = 11,
  IOCPUIDLE = 12,
  FSMBFREE = 13,
  FSINODEFREE = 14,
}
export enum ParserType { 
  UNDEF = 0,
  LOADAVERAGE = 1,
  CPU = 2,
  IO = 3,
  FS = 4,
  NET = 5,
}
