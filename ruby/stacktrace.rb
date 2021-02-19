require 'ddtrace'

def trigger_exception
    Datadog.tracer.trace("call.fetch") do
        v = {}.fetch("a")
    end
rescue => e
    raise ArgumentError.new("failed in trigger_exception")
end

def second_function
    Datadog.tracer.trace("call.trigger") do
        trigger_exception()
    end
rescue => e
    raise ArgumentError.new("failed in second_function")
end

def first_function
    Datadog.tracer.trace("call.second") do
        second_function()
    end
end

def main
    Datadog.tracer.trace("call.root") do
        first_function()
    end
end

main()