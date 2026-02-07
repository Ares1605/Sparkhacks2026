export default class API {
	constructor(base_url) {
		this.base_url = base_url
	}

	resolveImagePath(relpath) {
		return new URL(relpath, this.base_url).toString()
	}

	/**
	 * @returns {Promise<{ data: unknown } | { error: string }>}
	 */
	async getRecommendations() {
		const url = new URL("/recommendations", this.base_url)
		return await fetch(url.toString())
			.then(res => res.json())
			.catch(res => res.json())
	}
}
