from pyspark import SparkContext
import random

def inside(p):
    x, y = random.random(), random.random()
    return x*x + y*y < 1

sc = SparkContext.getOrCreate()
count = sc.parallelize(range(0, 100000)).filter(inside).count()
print("Pi is roughly %f" % (4.0 * count / 100000))

