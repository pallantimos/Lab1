import logging
import hashlib


def check_register(login, pass, pass2):
    hash_object = hashlib.md5(pass.encode('utf-8'))
    print(hash_object.hexdigest())


x, y = 3, 0
logging.basicConfig(level=logging.DEBUG, filename="py_log.log", filemode="w",
                    format="%(asctime)s %(levelname)s %(message)s")
logging.debug('logger configured')
logging.info('App Start')

check_register('aldar', '123', '123')

try:
    x/y
    logging.info(f"x/y successful with result: {x/y}.")
except ZeroDivisionError as err:
    logging.error("ZeroDivisionError",exc_info=True)