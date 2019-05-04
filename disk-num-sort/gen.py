import random
import struct
import argparse
import sys
import logging

FORMAT = '[%(asctime)-15s] %(message)s'
logging.basicConfig(format=FORMAT)
logger = logging.getLogger('sort')
logger.setLevel(logging.INFO)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('out', help='output file path')
    args = parser.parse_args(sys.argv[1:])

    logger.info('Generating random integers...')
    with open(args.out, 'wb') as f:
        n = int(10e7)
        step = int(10e4)
        batches = n // step
        for batch in range(batches):
            for i in random.sample(range(batch * step, step * (batch + 1)), step):
                # pack the integer to unsign in with big endian
                f.write(struct.pack('>I', i))
            f.flush()
            logger.info('Progress: {:.0f}%'.format(batch / batches * 100))
    logger.info('Done')
