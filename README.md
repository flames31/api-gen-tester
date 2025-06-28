# API Gen Tester CLI

A powerful Go-based CLI tool for generating, executing, and validating API test cases using large language models (LLMs) like OpenAI or Groq.

This tool helps developers and QA engineers automate edge-case generation for APIs and validate them efficiently using parallelized HTTP clients with retry logic and progress tracking.

---

## 🔧 Project Goal

- Automate the creation of diverse API test cases using LLMs based on a few user-provided samples.
- Run tests against a given API base URL with resilience and concurrency.
- Validate responses, track progress, and write results back to a JSON file.

---

## Quick Start Guide

1. Clone the repository

```bash
git clone https://github.com/your-username/api-gen-tester
cd api-gen-tester
```

2. Build the binary

```bash
go build -o api-gen-tester .
```

3. Provide your sample test cases in a sample.json file

4. Set your Groq Api-key in your .env file as
```bash
GROQ_APIKEY=<your_groq_api_key>
```

5. Generate and run new test cases powered by LLMs and view the results in the terminal and in the results.json file

```bash
./api-gen-tester generate --file sample.json
```

6. Optionally you can build a docker image and run it using the below commands

```bash
docker build -t api-gen-tester .
docker run -p 8080:8080 api-gen-tester generate --file sample.json
```