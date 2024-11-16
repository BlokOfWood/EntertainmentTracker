import api from './api';
import type { ApiResponse, CreateWorkRequest, UpdateWorkRequest, Work } from './api.model';

export async function getWorks(): ApiResponse<{ mediaEntries: Work[] }> {
	return api.get('/mediaentries');
}

export async function createWork(work: CreateWorkRequest): ApiResponse<Work> {
	return api.post('/mediaentries', work);
}

export async function getWork(id: number): ApiResponse<Work> {
	return api.get(`/mediaentries/${id}`);
}

export async function updateWork(id: number, work: UpdateWorkRequest): ApiResponse<Work> {
	return api.patch(`/mediaentries/${id}`, work);
}

export async function deleteWork(id: number): ApiResponse<void> {
	return api.delete(`/mediaentries/${id}`);
}