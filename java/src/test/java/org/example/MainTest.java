package org.example;

import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.assertEquals;


public class MainTest {
    @Test
    public void greet_NoParams_GreetsDojoers() {
        var expectedResult = "Hello Dojoers!";

        var result = Main.greet();

        assertEquals(expectedResult, result);
    }
}
