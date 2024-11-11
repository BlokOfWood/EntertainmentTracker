import api from './api';
import type { ApiResponse, Work } from './api.model';

export async function getWorks(): Promise<ApiResponse<{ mediaEntries: Work[] }>> {
	return api.get('/mediaentries');
}
