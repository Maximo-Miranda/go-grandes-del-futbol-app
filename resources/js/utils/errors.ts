export function fieldErrors(
    field: string | Record<string, string> | undefined,
): string[] {
    if (!field) return [];
    if (typeof field === "string") return [field];
    return Object.values(field);
}
