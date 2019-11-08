#pragma once

#include <sys/socket.h>

#include <memory>
#include <string>

namespace pl {
namespace stirling {
namespace testing {

// Convenience functions and predefined data for generating events expected from BPF socket probes.
class EventsFixture : public ::testing::Test {
 protected:
  static constexpr uint32_t kPID = 12345;
  static constexpr uint32_t kFD = 3;
  static constexpr uint64_t kPIDStartTimeTicks = 112358;

  void SetUp() {
    generation_ = 0;
    send_seq_num_ = 0;
    recv_seq_num_ = 0;
    current_ts_ns_ = 0;
  }

  struct socket_control_event_t InitConn(TrafficProtocol protocol) {
    struct socket_control_event_t conn_event {};
    conn_event.type = kConnOpen;
    conn_event.open.timestamp_ns = ++current_ts_ns_;
    conn_event.open.conn_id.pid = kPID;
    conn_event.open.conn_id.fd = kFD;
    conn_event.open.conn_id.generation = ++generation_;
    conn_event.open.conn_id.pid_start_time_ticks = kPIDStartTimeTicks;
    conn_event.open.addr.sin6_family = AF_INET;
    conn_event.open.traffic_class.protocol = protocol;
    conn_event.open.traffic_class.role = kRoleRequestor;
    return conn_event;
  }

  std::unique_ptr<SocketDataEvent> InitSendEvent(std::string_view msg) {
    return InitDataEvent(TrafficDirection::kEgress, TrafficProtocol::kProtocolHTTP, send_seq_num_++,
                         msg);
  }

  std::unique_ptr<SocketDataEvent> InitRecvEvent(std::string_view msg) {
    return InitDataEvent(TrafficDirection::kIngress, TrafficProtocol::kProtocolHTTP,
                         recv_seq_num_++, msg);
  }

  std::unique_ptr<SocketDataEvent> InitHTTP2SendEvent(std::string_view msg) {
    return InitDataEvent(TrafficDirection::kEgress, TrafficProtocol::kProtocolHTTP2,
                         send_seq_num_++, msg);
  }

  std::unique_ptr<SocketDataEvent> InitHTTP2RecvEvent(std::string_view msg) {
    return InitDataEvent(TrafficDirection::kIngress, TrafficProtocol::kProtocolHTTP2,
                         recv_seq_num_++, msg);
  }

  std::unique_ptr<SocketDataEvent> InitDataEvent(TrafficDirection direction,
                                                 TrafficProtocol protocol, uint64_t seq_num,
                                                 std::string_view msg) {
    socket_data_event_t event = {};
    event.attr.direction = direction;
    event.attr.traffic_class.protocol = protocol;
    event.attr.traffic_class.role = kRoleRequestor;
    event.attr.return_timestamp_ns = ++current_ts_ns_;
    event.attr.conn_id.pid = kPID;
    event.attr.conn_id.fd = kFD;
    event.attr.conn_id.generation = generation_;
    event.attr.conn_id.pid_start_time_ticks = kPIDStartTimeTicks;
    event.attr.seq_num = seq_num;
    event.attr.msg_size = msg.size();
    msg.copy(event.msg, msg.size());
    return std::make_unique<SocketDataEvent>(&event);
  }

  socket_control_event_t InitClose() {
    struct socket_control_event_t close_event {};
    close_event.type = kConnClose;
    close_event.close.timestamp_ns = ++current_ts_ns_;
    close_event.close.conn_id.pid = kPID;
    close_event.close.conn_id.fd = kFD;
    close_event.close.conn_id.generation = generation_;
    close_event.close.conn_id.pid_start_time_ticks = kPIDStartTimeTicks;
    close_event.close.rd_seq_num = recv_seq_num_;
    close_event.close.wr_seq_num = send_seq_num_;
    return close_event;
  }

  uint32_t generation_ = 0;
  uint64_t send_seq_num_ = 0;
  uint64_t recv_seq_num_ = 0;
  uint64_t current_ts_ns_ = 0;

  const std::string kHTTPReq0 =
      "GET /index.html HTTP/1.1\r\n"
      "Host: www.pixielabs.ai\r\n"
      "User-Agent: Mozilla/5.0 (X11; Linux x86_64)\r\n"
      "\r\n";

  const std::string kHTTPResp0 =
      "HTTP/1.1 200 OK\r\n"
      "Content-Type: application/json; charset=utf-8\r\n"
      "Content-Length: 5\r\n"
      "\r\n"
      "pixie";

  const std::string kHTTPReq1 =
      "GET /foo.html HTTP/1.1\r\n"
      "Host: www.pixielabs.ai\r\n"
      "User-Agent: Mozilla/5.0 (X11; Linux x86_64)\r\n"
      "\r\n";

  const std::string kHTTPResp1 =
      "HTTP/1.1 200 OK\r\n"
      "Content-Type: application/json; charset=utf-8\r\n"
      "Content-Length: 3\r\n"
      "\r\n"
      "foo";

  const std::string kHTTPReq2 =
      "GET /bar.html HTTP/1.1\r\n"
      "Host: www.pixielabs.ai\r\n"
      "User-Agent: Mozilla/5.0 (X11; Linux x86_64)\r\n"
      "\r\n";

  const std::string kHTTPResp2 =
      "HTTP/1.1 200 OK\r\n"
      "Content-Type: application/json; charset=utf-8\r\n"
      "Content-Length: 3\r\n"
      "\r\n"
      "bar";

  const std::string kHTTPUpgradeReq =
      "GET /index.html HTTP/1.1\r\n"
      "Host: www.pixielabs.ai\r\n"
      "Connection: Upgrade\r\n"
      "Upgrade: websocket\r\n"
      "\r\n";

  const std::string kHTTPUpgradeResp =
      "HTTP/1.1 101 Switching Protocols\r\n"
      "Upgrade: websocket\r\n"
      "Connection: Upgrade\r\n"
      "\r\n";

  static constexpr std::string_view kHTTP2EndStreamHeadersFrame =
      ConstStringView("\x0\x0\x0\x1\x5\x0\x0\x0\x1");
  static constexpr std::string_view kHTTP2EndStreamDataFrame =
      ConstStringView("\x0\x0\x0\x0\x1\x0\x0\x0\x1");
};

SocketDataEvent DataEventWithTimeSpan(std::string_view msg, TimeSpan time_span) {
  SocketDataEvent event;
  event.attr.entry_timestamp_ns = time_span.begin_ns;
  event.attr.return_timestamp_ns = time_span.end_ns;
  event.attr.msg_size = msg.size();
  event.msg = msg;
  return event;
}

}  // namespace testing
}  // namespace stirling
}  // namespace pl
