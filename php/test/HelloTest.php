<?php

use Emartech\CodingDojoStarter\Hello;
use PHPUnit\Framework\TestCase;

final class HelloTest extends TestCase
{
    public function test_greet_NoParameters_WelcomesDojoer(): void
    {
        $expectedResult = 'Hello, Dojoer!';

        $result = Hello::greet();

        $this->assertEquals($expectedResult, $result);
    }
}