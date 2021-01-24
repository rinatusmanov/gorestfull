package types

// IGormDB имитирует поведение *gorm.DB
type IGormDB interface {
	// возвращает экземпляр *gorm.DB
	GormDB() interface{}
	// Limit имитирует поведение Limit сущности *gorm.DB
	Limit(i int) IGormDB
	// Offset имитирует поведение Offset сущности *gorm.DB
	Offset(i int) IGormDB
	// Count имитирует поведение Count сущности *gorm.DB
	Count(count *int64) IGormDB
	// Unscoped имитирует поведение Unscoped сущности *gorm.DB
	Unscoped() IGormDB
	// Create имитирует поведение Create сущности *gorm.DB
	Create(value interface{}) IGormDB
	// Save имитирует поведение Save сущности *gorm.DB
	Save(value interface{}) IGormDB
	// Where имитирует поведение Where сущности *gorm.DB
	Where(query interface{}, args ...interface{}) IGormDB
	// Preload имитирует поведение Preload сущности *gorm.DB
	Preload(query string, args ...interface{}) IGormDB
	// Find имитирует поведение Find сущности *gorm.DB
	Find(dest interface{}, conds ...interface{}) IGormDB
	// Delete имитирует поведение Delete сущности *gorm.DB
	Delete(value interface{}, conds ...interface{}) IGormDB
	// Debug имитирует поведение Debug сущности *gorm.DB
	Debug() IGormDB
	// Error возвращает значение поля Error сущности *gorm.DB которую можно получить из GormDB
	Error() error
	// RowsAffected возвращает значение поля RowsAffected сущности *gorm.DB которую можно получить из GormDB
	RowsAffected() int64
}
