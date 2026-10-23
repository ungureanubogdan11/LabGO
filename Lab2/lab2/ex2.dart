import 'package:flutter/material.dart';

void main() {
  runApp(
    MaterialApp( // Removed 'const' here
      home: Scaffold( // Removed 'const' here
        body: SafeArea(
          child: Padding(
            padding: const EdgeInsets.all(16),
            child: const CarteDeVizita(), // const is fine here because constructor is const
          ),
        ),
      ),
    ),
  );
}

class CarteDeVizita extends StatelessWidget {
  const CarteDeVizita({super.key});

  @override
  Widget build(BuildContext context) {
    return Column(
      children: [
        const TextField(
          decoration: InputDecoration(hintText: "Ana Popescu"),
        ),
        const TextField(
          decoration: InputDecoration(hintText: "Student la FMI"),
        ),
        ElevatedButton(
          onPressed: () {},
          child: const Text("Trimite mesaj"),
        ),
      ],
    );
  }
}