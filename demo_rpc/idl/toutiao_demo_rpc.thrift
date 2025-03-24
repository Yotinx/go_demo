include "base.thrift"


// namespace & packageName
namespace go toutiao.demo.rpc

struct Item {
    1: required i64 id,
    2: required string title,
    3: required string content,
}

struct GetItemRequest {
    1: required i64 id,

    255: optional base.Base Base, // for compatibility with non-kitex framework
}

struct GetItemResponse {
    1: required Item item,

    255: optional base.BaseResp BaseResp, // for compatibility with non-kitex framework
}

//data struct + service


service ItemService {
    GetItemResponse GetItem(1: GetItemRequest req),
}
