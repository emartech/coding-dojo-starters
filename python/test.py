import unittest

from app import Hello

class TestHello(unittest.TestCase):

    def test_hello(self):
        expected_string = "Hello Dojoers!"

        result = Hello.greet()

        self.assertEqual(expected_string, result)


if __name__ == '__main__':
    unittest.main()