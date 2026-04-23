import client from './client';
import type { ApiResponse, Review, ReviewListResponse } from '../types';

export const reviewsApi = {
  getByDeviceId: (deviceId: string, page = 1, limit = 5) =>
    client.get<ApiResponse<ReviewListResponse>>(`/api/v1/reviews/device/${deviceId}`, { params: { page, limit } }),
  setRating: (deviceId: string, rating: number) =>
    client.post<ApiResponse<Review>>(`/api/v1/reviews/device/${deviceId}/rating`, { rating }),
  setComment: (deviceId: string, comment: string) =>
    client.post<ApiResponse<Review>>(`/api/v1/reviews/device/${deviceId}/comment`, { comment }),
  updateComment: (reviewId: string, comment: string) =>
    client.put<ApiResponse<Review>>(`/api/v1/reviews/${reviewId}/comment`, { comment }),
  remove: (reviewId: string) =>
    client.delete<ApiResponse<null>>(`/api/v1/reviews/${reviewId}`),
};

