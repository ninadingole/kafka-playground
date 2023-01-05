import org.jetbrains.kotlin.gradle.tasks.KotlinCompile

plugins {
    `kotlin-dsl`
    kotlin("jvm") version "1.7.10"
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation("org.jetbrains.kotlin:kotlin-stdlib-jdk8")

    api("junit:junit:4.13")
    implementation("junit:junit:4.13")
    testImplementation("junit:junit:4.13")
    testImplementation(kotlin("test"))
    compileOnly("org.apache.kafka:kafka-clients:2.8.0")
    compileOnly("org.apache.kafka:kafka-streams:2.8.0")
    compileOnly("org.apache.kafka:kafka-streams-test-utils:2.8.0")
    compileOnly("org.apache.kafka:kafka_2.13:2.8.0")
    implementation("com.google.code.gson:gson:2.8.5")

}

tasks.test {
    useJUnitPlatform()
}

tasks.withType<KotlinCompile> {
    kotlinOptions.jvmTarget = "1.8"
}

java {
    sourceCompatibility = JavaVersion.VERSION_1_8
    targetCompatibility = JavaVersion.VERSION_1_8
}

