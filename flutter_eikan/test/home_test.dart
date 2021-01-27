import 'package:flutter/material.dart';
import 'package:flutter_eikan/screens/home_screen.dart';
import 'package:flutter_test/flutter_test.dart';

Widget createHomeScreen() => MaterialApp(
      home: Scaffold(
        body: HomeScreen(),
      ),
    );
void main() {
  testWidgets('text find check', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(createHomeScreen());

    // Verify that our counter starts at 0.
    expect(find.text('HOGE'), findsOneWidget);
    expect(find.text('PIYO'), findsNothing);
  });
}
