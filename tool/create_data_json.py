import requests
import json
import time

import logging
logger = logging.getLogger()
logger.setLevel(logging.INFO)

API_ADDRESS = "https://r4v2wcv8c2.execute-api.ap-northeast-1.amazonaws.com/work/process2?type=main_summary:221309_hamamatsu_covid19_patients,main_summary:221309_hamamatsu_covid19_patients_summary,patients:221309_hamamatsu_covid19_patients,patients_summary:221309_hamamatsu_covid19_patients,inspection_persons:221309_hamamatsu_covid19_test_people,contacts:221309_hamamatsu_covid19_call_center"
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
                break
            else:
                hasError = False
                break

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
