import { err, httpClient, ok, ResponsePayload, Result } from "./common";

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
