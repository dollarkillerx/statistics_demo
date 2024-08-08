import { err, httpClient, ok,  } from './common';
import type {ResponsePayload, Result} from './common';

export interface TaskAccountItem {
  id: string
  created_at: string
  updated_at: string
  deleted_at: any
  client_id: string
  account: number
  leverage: number
  server: string
  company: string
  balance: number
  profit: number
  margin: number
}

// 获取所有账号
export const TaskAccount = async (): Promise<Result<TaskAccountItem[], string>> => {
  try {
    const resp = await httpClient.get('/api/accounts');

    // 将 resp 转换为 ResponsePayload 类型
    const value = resp as ResponsePayload<TaskAccountItem[]>;

    if (value.code !== 200) {
      console.log(value);
      return err(value.msg || "Unknown error");
    }

    if (!value.data) {
      return err("No data received");
    }

    return ok(value.data);
  } catch (error) {
    console.log(error);
    return err("Server request failed");
  }
}

export interface Task {
  positions: Position[]
  profits: Profit[]
}

export interface Position {
  id: string
  created_at: string
  updated_at: string
  deleted_at: any
  client_id: string
  order_id: number
  direction: string
  symbol: string
  magic: number
  open_price: number
  volume: number
  market: number
  swap: number
  profit: number
  common: string
  opening_time: number
  closing_time: number
  common_internal: string
  opening_time_system: number
  closing_time_system: number
}

export interface Profit {
  period: string
  max_profit: number
  min_profit: number
}

// 获取指定账号的任务
export const GetTaskByAccount = async (account: String): Promise<Result<Task, string>> => {
  try {
    const resp = await httpClient.get(`/api/account/${account}`);

    // 将 resp 转换为 ResponsePayload 类型
    const value = resp as ResponsePayload<Task>;

    if (value.code !== 200) {
      console.log(value);
      return err(value.msg || "Unknown error");
    }

    if (!value.data) {
      return err("No data received");
    }

    return ok(value.data);
  } catch (error) {
    console.log(error);
    return err("Server request failed");
  }
}
