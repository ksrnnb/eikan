import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:http/http.dart' as http;

class HomeScreen extends StatefulWidget {
  @override
  _HomeScreenState createState() => _HomeScreenState();
}

class _HomeScreenState extends State<HomeScreen> {
  String tmp = "test";
  String baseApiUrl = env['API_URL'];

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
              var response;
              try {
                response = await http.get(baseApiUrl);
                if (response.statusCode == 200) {
                  var data = jsonDecode(response.body);
                  setState(() {
                    tmp = data["message"];
                  });
                } else {
                  print("ERROR");
                }
              } catch (e) {
                print("URL is incorrect.");
                print(e);
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
