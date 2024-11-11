import api from './api';
import type { ApiResponse, CreateWorkRequest, UpdateWorkRequest, Work } from './api.model';

export async function getWorks(): Promise<ApiResponse<{ mediaEntries: Work[] }>> {
	return api.get('/mediaentries');
}

export async function createWork(work: CreateWorkRequest): Promise<ApiResponse<Work>> {
	return api.post('/mediaentries', work);
}

export async function getWork(id: number): Promise<ApiResponse<Work>> {
	return api.get(`/mediaentries/${id}`);
}

export async function updateWork(id: number, work: UpdateWorkRequest): Promise<ApiResponse<Work>> {
	return api.patch(`/mediaentries/${id}`, work);
}

export async function deleteWork(id: number): Promise<ApiResponse<void>> {
	return api.delete(`/mediaentries/${id}`);
}