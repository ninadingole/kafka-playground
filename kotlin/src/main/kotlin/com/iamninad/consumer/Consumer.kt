package com.iamninad.consumer

import com.google.gson.Gson
import org.apache.kafka.clients.consumer.KafkaConsumer
import java.time.Duration


fun main() {
    val inputTopic = "streams-persons-input"

    val props = mapOf(
        "bootstrap.servers" to "localhost:9092",
        "group.id" to "streams-persons-consumer",
        "key.deserializer" to "org.apache.kafka.common.serialization.StringDeserializer",
        "value.deserializer" to "org.apache.kafka.common.serialization.StringDeserializer"
    ).toProperties()

    val consumer = KafkaConsumer<String, String>(props).apply {
        subscribe(listOf(inputTopic))
    }

    consumer.use {
        while (true) {
            consumer
                .poll(Duration.ofMillis(100))
                .map { jsonToPerson(it.value()) }
                .forEach { println(it) }
        }
    }
}

fun jsonToPerson(value: String?): Person? {
    return Gson().fromJson(value, Person::class.java)
}

data class Person(val name: String, val age: Int)

