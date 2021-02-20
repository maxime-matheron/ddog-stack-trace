<?php
function trigger_exception() {
    try {
        throw new Exception('manually generated exception');
    } catch (Exception $e) {
        throw new Exception('failed in trigger_exception');
    }
}

function second_function() {
    try {
        trigger_exception();
    } catch (Exception $e) {
        throw new Exception('failed in second_function');
    }
}

function first_function() {
    second_function();
}

first_function()
?>