import 'dart:convert';
import 'package:http/http.dart' as http;

Future<void> main() async {
  const baseUrl = 'https://jsonplaceholder.typicode.com';

  final postResponse = await http.get(Uri.parse('$baseUrl/posts/1'));
  if (postResponse.statusCode == 200) {
    final postData = jsonDecode(postResponse.body);
    print(postData['title']);
  }

  final userResponse = await http.get(Uri.parse('$baseUrl/users/1'));
  if (userResponse.statusCode == 200) {
    final userData = jsonDecode(userResponse.body);
    print(userData['name']);
    print(userData['email']);
  }

  final missingResponse = await http.get(Uri.parse('$baseUrl/posts/99999'));
  print('Status: ${missingResponse.statusCode}');
}