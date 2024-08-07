import 'package:client/common/enums.dart';
import 'package:client/data/models/task.dart';
import 'package:client/data/source/task_source.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

part 'detail_task_state.dart';

class DetailTaskCubit extends Cubit<DetailTaskState> {
  DetailTaskCubit() : super(DetailTaskState(null, RequestStatus.init));

  fetchDetailTask(int id) async {
    emit(DetailTaskState(null, RequestStatus.loading));
    Task? result = await TaskSource.findById(id);
    if (result == null) {
      emit(DetailTaskState(null, RequestStatus.failed));
    } else {
      emit(DetailTaskState(result, RequestStatus.success));
    }
  }
}
