import 'package:flutter/material.dart';
import 'package:flutter_eikan/screens/home_screen.dart';

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    return MaterialApp(
      debugShowCheckedModeBanner: false,
      // TODO: 環境変数に
      // https://pub.dev/packages/flutter_dotenv
      title: '栄養管理アプリ（仮）',
      theme: ThemeData(
        primarySwatch: Colors.blue,
      ),
      home: Scaffold(
        body: HomeScreen(),
      ),
    );
  }
}
