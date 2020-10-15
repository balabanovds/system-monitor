// package: metric
// file: metric_service.proto

var metric_service_pb = require("./metric_service_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Metrics = (function () {
  function Metrics() {}
  Metrics.serviceName = "metric.Metrics";
  return Metrics;
}());

Metrics.GetStream = {
  methodName: "GetStream",
  service: Metrics,
  requestStream: false,
  responseStream: true,
  requestType: metric_service_pb.Request,
  responseType: metric_service_pb.Metric
};

exports.Metrics = Metrics;

function MetricsClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

MetricsClient.prototype.getStream = function getStream(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Metrics.GetStream, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

exports.MetricsClient = MetricsClient;

