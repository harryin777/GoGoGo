namespace go snow

struct SnowRequest {
    1: i64 wanted
}

struct SnowResponse {
    1: i64 snowCount
}

service Snow {
    SnowResponse wanted(1: SnowRequest req)
}