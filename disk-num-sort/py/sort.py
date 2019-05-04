import struct
import argparse
import sys
import logging

FORMAT = '[%(asctime)-15s] %(message)s'
logging.basicConfig(format=FORMAT)
logger = logging.getLogger('sort')
logger.setLevel(logging.DEBUG)


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument('infile', help='input random numbers file path')
    parser.add_argument('outfile', help='output sorted numbers file path')
    args = parser.parse_args(sys.argv[1:])

    num_mem = [0 for i in range(int(10e7))]

    logger.info('Sorting numbers...')
    with open(args.infile, 'rb') as infile:
        with open(args.outfile, 'wb') as outfile:
            # map each number in array with corresponding index
            while True:
                chunk = infile.read(4)
                if chunk == b'':
                    break
                try:
                    i, = struct.unpack('>I', chunk)
                except Exception as e:
                    logger.exception('failed to unpack: {}'.format(chunk))
                    raise e
                if num_mem[i] == 1:
                    raise ValueError('Duplicate number found: {}'.format(i))
                num_mem[i] = 1

            logger.info('Mapping Finished, begin to output result...')

            # output the result
            for idx, i in enumerate(num_mem):
                if (idx + 1) % 1000 == 0:
                    outfile.flush()
                if i == 1:
                    outfile.write(struct.pack('>I', idx))

    logger.info('Done')
