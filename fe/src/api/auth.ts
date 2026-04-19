import client from './client';
import type {
  ApiResponse,
  AuthResponse,
  LoginRequest,
  RegisterRequest,
  ChangePasswordRequest,
} from '../types';

export const authApi = {
  login: (data: LoginRequest) =>
    client.post<ApiResponse<AuthResponse>>('/auth/login', data),

  register: (data: RegisterRequest) =>
    client.post<ApiResponse>('/auth/register', data),

  changePassword: (data: ChangePasswordRequest) =>
    client.post<ApiResponse>('/auth/password/change', data),

  refresh: () =>
    client.post<ApiResponse<{ access_token: string }>>('/auth/refresh'),

  logout: () =>
    client.post<ApiResponse>('/auth/logout'),
};
