import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:http/http.dart' as http;

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  var tmp = "test";

  @override
  Widget build(BuildContext context) {
    return Center(
      child: Column(
        mainAxisAlignment: MainAxisAlignment.center,
        children: [
          Text(
            "HOGE",
            style: TextStyle(
              fontSize: 30,
              color: Colors.black,
            ),
          ),
          RaisedButton(
            onPressed: () async {
              // TODO: 環境変数
              // flutterはホスト側であることに注意。
              var response = await http.get('http://localhost:8000/v1');
              if (response.statusCode == 200) {
                var data = jsonDecode(response.body);
                setState(() {
                  tmp = data["message"];
                });
              } else {
                print("ERROR");
              }
            },
            child: Text(
              'Click me',
              style: TextStyle(
                fontSize: 20,
              ),
            ),
          ),
          SizedBox(
            height: 50,
          ),
          Text(
            tmp,
            style: TextStyle(
              fontSize: 30,
              color: Colors.black,
            ),
          )
        ],
      ),
    );
  }
}
