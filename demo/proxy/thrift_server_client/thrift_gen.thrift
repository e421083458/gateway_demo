namespace go thrift_gen

struct Data {
    1: string text
}

service format_data {
    Data do_format(1:Data data),
}

