// 通用返回
export interface ResponsePayload<T> {
  code: number;
  msg?: string;
  data?: T;
}

// 异常处理
export type Option<T> = Some<T> | null;

export class Some<T> {
  constructor(public value: T) {}
}

export type Result<T, E> = Ok<T> | Err<E>;

export class Ok<T> {
  constructor(public value: T) {}
}

export class Err<E> {
  constructor(public error: E) {}
}

// 工厂函数用于创建 Ok 和 Err 实例
export function ok<T>(value: T): Result<T, never> {
  return new Ok(value);
}

export function err<E>(error: E): Result<never, E> {
  return new Err(error);
}

/**
 * 错误处理演示
 * function division(a: number, b: number): Result<number, string> {
 *     if (b === 0) {
 *         return err('除数不能为0')
 *     } else {
 *         return ok(a / b)
 *     }
 * }
 *
 * function main() {
 *     const result = division(10, 0)
 *     if (result instanceof Ok) {
 *         console.log(result.value)
 *     } else {
 *         console.log(result.error)
 *     }
 * }
 */

// http

export interface FetchOptions extends RequestInit {
  url: string;
  method?: 'GET' | 'POST';
  headers?: HeadersInit;
  body?: any;
}

class HttpClient {
  private baseUrl: string;
  private headers: HeadersInit;

  constructor(baseUrl: string, headers?: HeadersInit) {
    this.baseUrl = baseUrl;
    this.headers = headers || {};
  }

  private async fetchRequest(options: FetchOptions): Promise<any> {
    const { url, method = 'GET', headers, body } = options;

    // 设置统一的headers
    const fetchHeaders: HeadersInit = {
      ...this.headers,
      ...headers,
    };

    // 请求拦截
    this.requestInterceptor(fetchHeaders);

    // 发送请求
    const response = await fetch(`${this.baseUrl}${url}`, {
      method,
      headers: fetchHeaders,
      body: body ? JSON.stringify(body) : null,
    });

    // 响应拦截
    this.responseInterceptor(response);

    if (!response.ok) {
      throw new Error(`HTTP error! status: ${response.status}`);
    }

    return response.json();
  }

  public get(url: string, headers?: HeadersInit): Promise<any> {
    return this.fetchRequest({ url, method: 'GET', headers });
  }

  public post(url: string, body: any, headers?: HeadersInit): Promise<any> {
    return this.fetchRequest({ url, method: 'POST', headers, body });
  }

  private requestInterceptor(headers: HeadersInit): void {
    // 在这里可以添加统一的请求拦截逻辑
    console.log('Request Interceptor:', headers);
  }

  private responseInterceptor(response: Response): void {
    // 在这里可以添加统一的响应拦截逻辑
    console.log('Response Interceptor:', response);
  }
}

export const httpClient = new HttpClient('', {
  'Content-Type': 'application/json',
  // 'Authorization': 'Bearer YOUR_TOKEN_HERE',
});

// 192.168.40.238
