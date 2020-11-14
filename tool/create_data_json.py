import requests
import json
import time

import logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

API_ADDRESS = "https://r4v2wcv8c2.execute-api.ap-northeast-1.amazonaws.com/work/process?type=main_summary:5ab47071-3651-457c-ae2b-bfb8fdbe1af1,main_summary:92f9ebcd-a3f1-4d5d-899b-d69214294a45,patients:5ab47071-3651-457c-ae2b-bfb8fdbe1af1,patients_summary:5ab47071-3651-457c-ae2b-bfb8fdbe1af1,inspection_persons:d4827176-d887-412a-9344-f84f161786a2,contacts:1b57f2c0-081e-4664-ba28-9cce56d0b314"
RETRY_COUNT = 3
RETRY_WAIT_SEC = 3

def process(apiKey):
    try:
        apiResponse = requestWithRetry(apiKey)

        return apiResponse.text

    except Exception as e:
        logger.exception(e)
        raise e

def requestWithRetry(apiKey):

    hasError = False
    apiResponse = None
    index = 0
    for index in range(RETRY_COUNT):
        try:
            apiResponse = requests.get(API_ADDRESS, headers={"x-api-key": apiKey})
            di = json.loads(apiResponse.text)
            if(di["hasError"]):
                hasError = True
                break;

        except Exception as e:
            if(index + 1 < RETRY_COUNT):
                logger.warn("request failed : loop count = {0}".format(index))
                logger.warn(e)
                time.sleep(RETRY_WAIT_SEC)
            else:
                logger.error("request failed : loop count = {0}".format(index))
                logger.exception(e)
                raise Exception(e)

    if(hasError):
        raise Exception("has error!")

    return apiResponse
