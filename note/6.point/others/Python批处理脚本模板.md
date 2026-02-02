```python
#!/usr/bin/env python3

import time
import queue
import csv
import logging

from datetime import datetime
from concurrent.futures import ThreadPoolExecutor


THREAD_POOL_MAX_WORKERS = 20 + 2
SLEEP_TIME = 0.3
SUCCESS_OUTPUT_CHANNEL = 0
FAIL_OUTPUT_CHANNEL = 1
INPUT_FILE = "addresses.csv"
SUCCESS_OUTPUT_FILE = "success.csv"
FAIL_OUTPUT_FILE = "fail.csv"
LOG_FILE = "log.log"
SUCCESS_CSV_HEADER = [
    "<TO BE OVERRIDDEN>",
]
FAIL_CSV_HEADER = [
    "<TO BE OVERRIDDEN>",
]


threadPool = ThreadPoolExecutor(max_workers=THREAD_POOL_MAX_WORKERS)
success_data_queue = queue.Queue()
fail_data_queue = queue.Queue()


def csv_writer_thread(channel):
    match channel:
        case c if c == SUCCESS_OUTPUT_CHANNEL:
            file_name = SUCCESS_OUTPUT_FILE
            queue_obj = success_data_queue
            header = SUCCESS_CSV_HEADER
        case c if c == FAIL_OUTPUT_CHANNEL:
            file_name = FAIL_OUTPUT_FILE
            queue_obj = fail_data_queue
            header = FAIL_CSV_HEADER
    with open(file_name, "w", newline="", encoding="utf-8") as file:
        writer = csv.DictWriter(file, fieldnames=header)
        writer.writeheader()
        while True:
            item = queue_obj.get()
            if item is None:
                break
            writer.writerow(item)
            queue_obj.task_done()


def process(input_row):
    output_channel = None
    try:
        output_channel = SUCCESS_OUTPUT_CHANNEL
        return output_channel
    except Exception as e:
        input_row["fail_reason"] = f"unexpected error: {str(e)}"
        output_channel = FAIL_OUTPUT_CHANNEL
        logging.error(f"[process] Unexpected error for record={input_row}, err: {e}")
        return output_channel
    finally:
        match output_channel:
            case c if c == SUCCESS_OUTPUT_CHANNEL:
                success_data_queue.put(input_row)
            case c if c == FAIL_OUTPUT_CHANNEL:
                fail_data_queue.put(input_row)
            case _:
                input_row["fail_reason"] = "unknown error"
                fail_data_queue.put(input_row)
        time.sleep(SLEEP_TIME)


def batch_process():
    try:
        threadPool.submit(csv_writer_thread, SUCCESS_OUTPUT_CHANNEL)
        threadPool.submit(csv_writer_thread, FAIL_OUTPUT_CHANNEL)

        with open(INPUT_FILE, "r", encoding="utf-8") as input_file:
            async_results = []
            reader = csv.DictReader(input_file)

            for row in reader:
                async_result = threadPool.submit(process, row)
                async_results.append(async_result)

            success_count = 0
            data_size = len(async_results)

            for async_result in tqdm(async_results, total=data_size, desc="Processing"):
                channel = async_result.result()
                match channel:
                    case c if c == SUCCESS_OUTPUT_CHANNEL:
                        success_count += 1

            print(
                f"total: {data_size}, success: {success_count}, failed: {data_size - success_count}"
            )
    finally:
        success_data_queue.put(None)
        fail_data_queue.put(None)


def main():
    logging.basicConfig(
        format="%(asctime)s - %(message)s", level=logging.ERROR, filename=LOG_FILE
    )
    requests_log = logging.getLogger("urllib3")
    requests_log.setLevel(logging.ERROR)
    requests_log.propagate = True

    logging.error("Beginning script...")
    batch_process()
    logging.error("Finished script.")


if __name__ == "__main__":
    start_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    main()
    end_time = datetime.now().strftime("%Y-%m-%d %H:%M:%S")
    print(f"Start time: {start_time}. End time: {end_time}")
    print("Execute finished, see output log and files to check result")

```

