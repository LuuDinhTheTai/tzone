import client from './client';
import type { AIChatRequest, AIChatResponse, AIVideoReviewRequest, AIVideoReviewResponse, ApiResponse } from '../types';

export const aiApi = {
  chat: (payload: AIChatRequest) =>
    client.post<ApiResponse<AIChatResponse>>('/api/v1/ai/chat', payload),
  videoReviews: (payload: AIVideoReviewRequest) =>
    client.post<ApiResponse<AIVideoReviewResponse>>('/api/v1/ai/video-reviews', payload),
};

