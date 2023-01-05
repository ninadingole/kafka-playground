package com.iamninad.streams

import org.apache.kafka.common.serialization.Serdes
import org.apache.kafka.streams.KafkaStreams
import org.apache.kafka.streams.StreamsBuilder
import org.apache.kafka.streams.kstream.Consumed

interface ErrorHandler {
    fun handle(e: Exception)
}

fun main(args: Array<String>) {
    var builder = StreamsBuilder()
    builder.stream("streams-persons-input", Consumed.with(Serdes.String(), Serdes.String()))
//            .mapValues { value -> jsonToPerson(value) }
            .to("streams-persons-output")

    var topology = builder.build()

    var streams = KafkaStreams(topology, mapOf(
        "application.id" to "streams-persons",
        "bootstrap.servers" to "localhost:9092",
        "default.key.serde" to Serdes.String().javaClass.name,
        "default.value.serde" to Serdes.String().javaClass.name
    ).toProperties())

    streams.start()
}

