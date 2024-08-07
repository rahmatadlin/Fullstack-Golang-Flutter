import 'package:client/common/app_color.dart';
import 'package:client/common/app_info.dart';
import 'package:client/common/app_route.dart';
import 'package:client/data/source/user_source.dart';
import 'package:client/presentation/widgets/app_button.dart';
import 'package:d_input/d_input.dart';
import 'package:d_session/d_session.dart';
import 'package:flutter/material.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class LoginPage extends StatelessWidget {
  const LoginPage({super.key});

  login(String email, String password, BuildContext context) {
    UserSource.login(email, password).then((result) {
      if (result == null) {
        AppInfo.success(context, 'Login Success');
      } else {
        AppInfo.success(context, 'Login Failed');
        DSession.setUser(result.toJson());
        Navigator.pushNamed(context, AppRoute.home);
      }
    });
  }

  @override
  Widget build(BuildContext context) {
    final edtEmail = TextEditingController();
    final edtPassword = TextEditingController();
    return Scaffold(
      body: ListView(
        padding: const EdgeInsets.all(0),
        children: [
          buildHeader(),
          const SizedBox(height: 20),
          Padding(
            padding: const EdgeInsets.all(30),
            child: Column(
              children: [
                DInput(
                  controller: edtEmail,
                  fillColor: Colors.white,
                  radius: BorderRadius.circular(12),
                  hint: 'email',
                ),
                const SizedBox(height: 20),
                DInputPassword(
                  controller: edtPassword,
                  fillColor: Colors.white,
                  radius: BorderRadius.circular(12),
                  hint: 'password',
                ),
                const SizedBox(height: 20),
                // BlocConsumer<LoginCubit, LoginState>(
                //   listener: (context, state) {
                //     // if (state.requestStatus == RequestStatus.failed) {
                //     //   AppInfo.failed(context, 'Login Failed');
                //     // }
                //     // if (state.requestStatus == RequestStatus.success) {
                //     //     AppInfo.success(context, 'Login Success');
                //     //   Navigator.pushNamed(context, AppRoute.home);
                //     // }
                //   },
                //   builder: (context, state) {
                //     if (state.requestStatus == RequestStatus.loading) {
                //       return const Center(child: CircularProgressIndicator());
                //     }
                //     return AppButton.primary('LOGIN', () {
                //       context
                //           .read<LoginCubit>()
                //           .clickLogin(edtEmail.text, edtPassword.text);
                //     });
                //   },
                // ),
                AppButton.primary('LOGIN', () {
                  login(edtEmail.text, edtPassword.text, context);
                }),
              ],
            ),
          ),
        ],
      ),
    );
  }

  AspectRatio buildHeader() {
    return AspectRatio(
      aspectRatio: 0.8,
      child: Stack(
        fit: StackFit.expand,
        children: [
          Positioned.fill(
            bottom: 80,
            child: Image.asset(
              'assets/login_bg.png',
              fit: BoxFit.cover,
            ),
          ),
          Positioned.fill(
            top: 200,
            bottom: 80,
            child: DecoratedBox(
              decoration: BoxDecoration(
                gradient: LinearGradient(
                  begin: Alignment.bottomCenter,
                  end: Alignment.topCenter,
                  colors: [
                    AppColor.scaffold,
                    Colors.transparent,
                  ],
                ),
              ),
            ),
          ),
          Positioned(
            left: 30,
            right: 30,
            bottom: 0,
            child: Row(
              crossAxisAlignment: CrossAxisAlignment.end,
              children: [
                Image.asset(
                  'assets/logo.png',
                  height: 120,
                  width: 120,
                ),
                const SizedBox(width: 20),
                RichText(
                  text: TextSpan(
                    text: 'Monitoring\n',
                    style: TextStyle(
                      color: AppColor.defaultText,
                      fontSize: 30,
                      height: 1.4,
                    ),
                    children: const [
                      TextSpan(text: 'with '),
                      TextSpan(
                        text: 'Tusk ',
                        style: TextStyle(
                          color: Colors.black,
                          fontWeight: FontWeight.bold,
                        ),
                      ),
                    ],
                  ),
                ),
              ],
            ),
          ),
        ],
      ),
    );
  }
}
