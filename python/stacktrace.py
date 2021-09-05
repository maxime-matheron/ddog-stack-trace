from ddtrace import tracer
from ddtrace.opentracer import Tracer, set_global_tracer

def init_tracer(service_name):
    config = {
      'agent_hostname': 'localhost',
      'agent_port': 8126,
    }
    tracer = Tracer(service_name, config=config)
    set_global_tracer(tracer)
    return tracer

init_tracer('stacktrace')

def fourth_function():
    try:
        with tracer.trace('run.operation'):
            v = {}['a']
    except KeyError as e:
        raise ValueError('failed in fourth_function')

def third_function():
    try:
        with tracer.trace('call.fourth_function'):
            fourth_function()
    except ValueError as e:
        raise ValueError('failed to call fourth_function in third_function')

def second_function():
    try:
        with tracer.trace('call.third_function'):
            third_function()
    except ValueError as e:
        raise ValueError('failed to call third_function in second_function') from e

def first_function():
    with tracer.trace('call.second_function'):
        second_function()

with tracer.trace('call.first_function'):
    first_function()
