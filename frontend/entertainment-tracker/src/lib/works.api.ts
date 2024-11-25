import api from './api';
import type {
	ApiResponse,
	CreateWorkRequest,
	SharedWork,
	ShareWorkRequest,
	UpdateWorkRequest,
	Work
} from './api.model';

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

export async function shareWork(req: ShareWorkRequest): ApiResponse<void> {
	return api.post('/share', req);
}

export async function getSharedWorks(): ApiResponse<{ sharedEntries: SharedWork[] }> {
	return api.get('/shared');
}
