import 'package:flutter/material.dart';
import 'package:flutter_eikan/screens/home_screen.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart' as DotEnv;

Widget createHomeScreen() => MaterialApp(
      home: Scaffold(
        body: HomeScreen(),
      ),
    );
Future main() async {
  // テスト実行時に.envファイルのパスを正しく取得するのに必要
  TestWidgetsFlutterBinding.ensureInitialized();
  await DotEnv.load();

  testWidgets('text find check', (WidgetTester tester) async {
    // Build our app and trigger a frame.
    await tester.pumpWidget(createHomeScreen());

    // Verify that our counter starts at 0.
    expect(find.text('HOGE'), findsOneWidget);
    expect(find.text('PIYO'), findsNothing);
  });
}
