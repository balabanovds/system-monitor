/**
 * @fileoverview gRPC-Web generated client stub for metric
 * @enhanceable
 * @public
 */

// GENERATED CODE -- DO NOT EDIT!


/* eslint-disable */
// @ts-nocheck



const grpc = {};
grpc.web = require('grpc-web');


var google_protobuf_timestamp_pb = require('google-protobuf/google/protobuf/timestamp_pb.js')

var google_protobuf_empty_pb = require('google-protobuf/google/protobuf/empty_pb.js')
const proto = {};
proto.metric = require('./metric_service_pb.js');

/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.metric.MetricsClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @param {string} hostname
 * @param {?Object} credentials
 * @param {?Object} options
 * @constructor
 * @struct
 * @final
 */
proto.metric.MetricsPromiseClient =
    function(hostname, credentials, options) {
  if (!options) options = {};
  options['format'] = 'text';

  /**
   * @private @const {!grpc.web.GrpcWebClientBase} The client
   */
  this.client_ = new grpc.web.GrpcWebClientBase(options);

  /**
   * @private @const {string} The hostname
   */
  this.hostname_ = hostname;

};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.metric.Request,
 *   !proto.metric.Metric>}
 */
const methodDescriptor_Metrics_GetStream = new grpc.web.MethodDescriptor(
  '/metric.Metrics/GetStream',
  grpc.web.MethodType.SERVER_STREAMING,
  proto.metric.Request,
  proto.metric.Metric,
  /**
   * @param {!proto.metric.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metric.Metric.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.metric.Request,
 *   !proto.metric.Metric>}
 */
const methodInfo_Metrics_GetStream = new grpc.web.AbstractClientBase.MethodInfo(
  proto.metric.Metric,
  /**
   * @param {!proto.metric.Request} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metric.Metric.deserializeBinary
);


/**
 * @param {!proto.metric.Request} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.metric.Metric>}
 *     The XHR Node Readable Stream
 */
proto.metric.MetricsClient.prototype.getStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/metric.Metrics/GetStream',
      request,
      metadata || {},
      methodDescriptor_Metrics_GetStream);
};


/**
 * @param {!proto.metric.Request} request The request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!grpc.web.ClientReadableStream<!proto.metric.Metric>}
 *     The XHR Node Readable Stream
 */
proto.metric.MetricsPromiseClient.prototype.getStream =
    function(request, metadata) {
  return this.client_.serverStreaming(this.hostname_ +
      '/metric.Metrics/GetStream',
      request,
      metadata || {},
      methodDescriptor_Metrics_GetStream);
};


/**
 * @const
 * @type {!grpc.web.MethodDescriptor<
 *   !proto.google.protobuf.Empty,
 *   !proto.metric.ParsersInfoResponse>}
 */
const methodDescriptor_Metrics_ParsersInfo = new grpc.web.MethodDescriptor(
  '/metric.Metrics/ParsersInfo',
  grpc.web.MethodType.UNARY,
  google_protobuf_empty_pb.Empty,
  proto.metric.ParsersInfoResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metric.ParsersInfoResponse.deserializeBinary
);


/**
 * @const
 * @type {!grpc.web.AbstractClientBase.MethodInfo<
 *   !proto.google.protobuf.Empty,
 *   !proto.metric.ParsersInfoResponse>}
 */
const methodInfo_Metrics_ParsersInfo = new grpc.web.AbstractClientBase.MethodInfo(
  proto.metric.ParsersInfoResponse,
  /**
   * @param {!proto.google.protobuf.Empty} request
   * @return {!Uint8Array}
   */
  function(request) {
    return request.serializeBinary();
  },
  proto.metric.ParsersInfoResponse.deserializeBinary
);


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @param {function(?grpc.web.Error, ?proto.metric.ParsersInfoResponse)}
 *     callback The callback function(error, response)
 * @return {!grpc.web.ClientReadableStream<!proto.metric.ParsersInfoResponse>|undefined}
 *     The XHR Node Readable Stream
 */
proto.metric.MetricsClient.prototype.parsersInfo =
    function(request, metadata, callback) {
  return this.client_.rpcCall(this.hostname_ +
      '/metric.Metrics/ParsersInfo',
      request,
      metadata || {},
      methodDescriptor_Metrics_ParsersInfo,
      callback);
};


/**
 * @param {!proto.google.protobuf.Empty} request The
 *     request proto
 * @param {?Object<string, string>} metadata User defined
 *     call metadata
 * @return {!Promise<!proto.metric.ParsersInfoResponse>}
 *     A native promise that resolves to the response
 */
proto.metric.MetricsPromiseClient.prototype.parsersInfo =
    function(request, metadata) {
  return this.client_.unaryCall(this.hostname_ +
      '/metric.Metrics/ParsersInfo',
      request,
      metadata || {},
      methodDescriptor_Metrics_ParsersInfo);
};


module.exports = proto.metric;

