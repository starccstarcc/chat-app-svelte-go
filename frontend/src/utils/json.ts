function formatJSON(data: string): { [key: string]: unknown }[] {
	const parseData = JSON.parse(data);

	if (Array.isArray(parseData))
		return parseData.map((value, index) => ({ key: `${index}`, value }));

	return Object.entries(parseData).map(([key, value]) => ({ key, value }));
}

function formatValuesIntoString(data: unknown): string {
	if (!data) return '';

	if (typeof data === 'string') return data;

	if (typeof data === 'object') {
		if (Array.isArray(data)) {
			return data.map((value: unknown) => formatValuesIntoString(value)).join('\n');
		} else {
			console.log(data);
			return Object.entries(data).reduce((acc, [key, value]) => {
				acc += `\n${key}: ${formatValuesIntoString(value)}\n\n`;

				return acc;
			}, '');
		}
	}

	return data.toString();
}

export { formatJSON, formatValuesIntoString };
