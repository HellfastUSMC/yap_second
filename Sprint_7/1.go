package main

//var dumpFile *os.File
//
//func init() {
//	if dumpFile == nil {
//		// создаём файл при первом обращении
//		fname, err := os.Executable()
//		if err == nil {
//			dumpFile, err = os.Create(fname + `.dump`)
//		}
//		if err != nil {
//			panic(err)
//		}
//
//	}
//}
//
//func Dump(data []byte) error {
//	_, err := dumpFile.Write(data)
//	return err
//}

//	type Config struct {
//		Version string
//		Plugins []string
//		Stat    map[string]int
//	}
//
//	func (cfg *Config) Clone() *Config {
//		clone := &Config{
//			Version: cfg.Version,
//			Plugins: make([]string, len(cfg.Plugins)),
//			Stat:    make(map[string]int, len(cfg.Stat)),
//		}
//		for i, val := range cfg.Plugins {
//			clone.Plugins[i] = val
//		}
//		for key, val := range cfg.Stat {
//			clone.Stat[key] = val
//		}
//		return clone
//	}
//type Car struct {
//	Type  string
//	Seats int // количество мест
//	Doors int // количество дверей
//	ABS   bool
//}
//
//// CarOptionFunc определяет тип функции для опций.
//type CarOptionFunc func(*Car)
//
//// NewCar создаёт автомобиль, используя указанные опции.
//func NewCar(cartype string, opts ...CarOptionFunc) *Car {
//	car := &Car{}
//	for _, fnc := range opts {
//		fnc(car)
//	}
//	return car
//}
//
//func (c *Car) String() string {
//	return fmt.Sprintf("%s [seats:%d / doors: %d / abs: %t]",
//		c.Type, c.Seats, c.Doors, c.ABS)
//}
//
//// WithSeats определяет количество мест в автомобиле.
//func WithSeats(seats int) CarOptionFunc {
//	return func(car *Car) {
//		car.Seats = seats
//	}
//}
//
//// WithDoors определяет количество дверей.
//func WithDoors(doors int) CarOptionFunc {
//	return func(car *Car) {
//		car.Doors = doors
//	}
//}
//
//// WithABS указывает на наличие ABS в автомобиле.
//func WithABS() CarOptionFunc {
//	return func(car *Car) {
//		car.ABS = true
//	}
//}
//
//func TestNewCar(t *testing.T) {
//	sportsCar := NewCar("sports car", WithSeats(2), WithDoors(2))
//	minivan := NewCar("minivan", WithSeats(7), WithDoors(5), WithABS())
//
//	if sportsCar.String() != "sports car [seats:2 / doors: 2 / abs: false]" {
//		t.Errorf("wrong %s", sportsCar.String())
//	}
//	if minivan.String() != "minivan [seats:7 / doors: 5 / abs: true]" {
//		t.Errorf("wrong %s", minivan.String())
//	}
//}
//
//type DatabaseConnector interface {
//	Query(q string) error
//}
//
//type MysqlConnector struct {
//}
//
//func newMysqlConnector(dsn string) *MysqlConnector {
//	fmt.Println("Connect to mysql")
//	return &MysqlConnector{}
//}
//
//func (c *MysqlConnector) Query(q string) error {
//	fmt.Printf("Query to mysql: %s\n", q)
//	return nil
//}
//
//type PostgresqlConnector struct {
//}
//
//func newPostgresqlConnector(dsn string) *PostgresqlConnector {
//	fmt.Println("Connect to postgresql")
//	return &PostgresqlConnector{}
//}
//
//func (c *PostgresqlConnector) Query(q string) error {
//	fmt.Printf("Query to postgresql: %s\n", q)
//	return nil
//}
//
//type SQLiteConnector struct {
//}
//
//func newSQLiteConnector(dsn string) *SQLiteConnector {
//	fmt.Println("Connect to sqlite")
//	return &SQLiteConnector{}
//}
//
//func (c *SQLiteConnector) Query(q string) error {
//	fmt.Printf("Query to sqlite: %s\n", q)
//	return nil
//}
//
//// NewConnector реализует фабричный метод.
//func NewConnector(dsn string) DatabaseConnector {
//	switch {
//	case strings.HasPrefix(dsn, "mysql://"):
//		return newMysqlConnector(dsn)
//	case strings.HasPrefix(dsn, "postgresql://"):
//		return newPostgresqlConnector(dsn)
//	case strings.HasPrefix(dsn, "sqlite://"):
//		return newSQLiteConnector(dsn)
//	default:
//		panic(fmt.Sprintf("unknown dsn protocol: %s", dsn))
//	}
//}
//
//func main() {
//	mysql := NewConnector("mysql://...")
//	mysql.Query("SELECT something FROM list")
//
//	pg := NewConnector("postgresql://...")
//	pg.Query("SELECT something FROM list")
//
//	sqlite := NewConnector("sqlite://...")
//	sqlite.Query("SELECT something FROM list")
//}

// Customer — покупатель.
//type Customer struct {
//	Name string
//}
//
//func (c *Customer) SetName(name string) {
//	c.Name = name
//}
//
//func (c *Customer) String() string {
//	return fmt.Sprintf("Customer: %s", c.Name)
//}
//
//// Seller — продавец.
//type Seller struct {
//	Name string
//}
//
//func (c *Seller) SetName(name string) {
//	c.Name = name
//}
//
//func (c *Seller) String() string {
//	return fmt.Sprintf("Seller: %s", c.Name)
//}
//
//// Provider — интерфейс фабрики.
//type Provider interface {
//	NewCustomer() *Customer
//	NewSeller() *Seller
//}
//
//type GoogleAuth struct {
//}
//
//func (g *GoogleAuth) NewCustomer() *Customer {
//	var customer Customer
//	// получаем имя из Google-аккаунта
//	name := "Google Customer"
//	customer.SetName(name)
//	return &customer
//}
//
//func (g *GoogleAuth) NewSeller() *Seller {
//	var seller Seller
//	// получаем имя из Google-аккаунта
//	name := "Google Seller"
//	seller.SetName(name)
//	return &seller
//}
//
//type YandexAuth struct {
//}
//
//func (g *YandexAuth) NewCustomer() *Customer {
//	var customer Customer
//	// получаем имя из Google-аккаунта
//	name := "Yandex Customer"
//	customer.SetName(name)
//	return &customer
//}
//
//func (g *YandexAuth) NewSeller() *Seller {
//	var seller Seller
//	// получаем имя из Google-аккаунта
//	name := "Yandex Seller"
//	seller.SetName(name)
//	return &seller
//}
//
//// допишите код
//// ...
//
//// AuthFactory — абстрактная фабрика аутентификации.
//func AuthFactory(provider string) Provider {
//	switch provider {
//	case "google":
//		return &GoogleAuth{}
//	case "yandex":
//		return &YandexAuth{}
//	default:
//		panic(fmt.Sprintf("unknown provider %s", provider))
//	}
//	return nil
//}
//
//func TestAuth(t *testing.T) {
//	googleAuth := AuthFactory("google")
//	yandexAuth := AuthFactory("yandex")
//
//	if googleAuth.NewCustomer().String() != "Customer: Google Customer" {
//		t.Errorf("wrong google customer")
//	}
//	if googleAuth.NewSeller().String() != "Seller: Google Seller" {
//		t.Errorf("wrong google seller")
//	}
//	if yandexAuth.NewCustomer().String() != "Customer: Yandex Customer" {
//		t.Errorf("wrong yandex customer")
//	}
//	if yandexAuth.NewSeller().String() != "Seller: Yandex Seller" {
//		t.Errorf("wrong yandex seller")
//	}
//}

//type Modifier interface {
//	Modify() string
//}
//
//type Original struct {
//	Value string
//}
//
//func (o *Original) Modify() string {
//	return o.Value
//}
//
//// Upper возвращает строку в верхнем регистре.
//type Upper struct {
//	modifier Modifier
//}
//
//func (u *Upper) Modify() string {
//	return strings.ToUpper(u.modifier.Modify())
//}
//
//// добавьте метод Modify для *Upper
//// он должен возвращать строку в верхнем регистре
//// ...
//
//// Replace заменяет строки old на new.
//type Replace struct {
//	modifier Modifier
//	old      string
//	new      string
//}
//
//// добавьте метод Modify для *Replace
//// он должен заменять old на new
//// ...
//
//func (r *Replace) Modify() string {
//	return strings.Replace(r.modifier.Modify(), r.old, r.new, 1)
//}
//
//func TestModifier(t *testing.T) {
//	original := &Original{Value: "Привет, гофер!"}
//	replace := &Replace{
//		// инициализируйте поля нужными значениями
//		// ...
//		modifier: original,
//		old:      "гофер",
//		new:      "мир",
//	}
//	upper := &Upper{
//		// инициализируйте поле нужным значением
//		// ...
//		modifier: replace,
//	}
//	if upper.Modify() != "ПРИВЕТ, МИР!" {
//		t.Errorf(`get %s`, upper.Modify())
//	}
//}

//// JSONData — интерфейс для декодирования JSON.
//type JSONData interface {
//	DecodeJSON() interface{}
//}
//
//// YAMLData — интерфейс для декодирования YAML.
//type YAMLData interface {
//	DecodeYAML() interface{}
//}
//
//type Client struct {
//	Data interface{}
//}
//
//func (client *Client) Decode(input JSONData) {
//	client.Data = input.DecodeJSON()
//}
//
//// добавьте тип Adapter и необходимый метод
//// ...
//type Adapter struct {
//	data YAMLData
//}
//
//func (a *Adapter) DecodeJSON() interface{} {
//	return a.data.DecodeYAML()
//}
//
//func Load(client Client, input YAMLData) {
//	adapter := &Adapter{
//		data: input,
//	}
//	client.Decode(adapter)
//}

//package main
//
//import (
//"fmt"
//"testing"
//)
//
//type Operation int
//
//const (
//	Add Operation = iota
//	Sub
//	Mul
//	Div
//)
//
//type Calculator interface {
//	Calculate() int
//}
//
//// допишите код
//// ...
//
//type Oper struct {
//	Type Operation
//	Left Calculator
//	Right Calculator
//}
//
//func (o *Oper) Calculate() int {
//	switch o.Type {
//	case Add:
//		return o.Left.Calculate() + o.Right.Calculate()
//
//	case Sub:
//		return o.Left.Calculate() - o.Right.Calculate()
//
//	case Mul:
//		return o.Left.Calculate() * o.Right.Calculate()
//
//	case Div:
//		return o.Left.Calculate() / o.Right.Calculate()
//	default:
//		return 0
//	}
//}
//
//type Number struct {
//	Value int
//}
//
//func (n *Number) Calculate() int{
//	return n.Value
//}
//
//func TestCalc(t *testing.T) {
//	root := &Oper{
//		Type: Div,
//		Left: &Oper{
//			Type: Mul,
//			Left: &Oper{
//				Type:  Add,
//				Left:  &Number{Value: 2},
//				Right: &Number{Value: 3},
//			},
//			Right: &Oper{
//				Type:  Sub,
//				Left:  &Number{Value: 77},
//				Right: &Number{Value: 55},
//			},
//		},
//		Right: &Number{Value: 2},
//	}
//	if root.Calculate() != 55 {
//		t.Errorf(`get %d want %d`, root.Calculate(), 77)
//	}
//}

//package main
//
//import (
//"fmt"
//"unicode"
//)
//
//type TokenType int
//
//const (
//	TNumber = iota
//)
//
//const (
//	StateMain = iota
//	StateNumber
//)
//
//// Token — информация о токене.
//type Token struct {
//	Type  TokenType
//	Value string
//}
//
//// State — интерфейс состояния. Next передаёт очередной символ.
//// Если символ разобран, то возвращается true.
//type State interface {
//	Next(rune) bool
//}
//
//// Number определяет число.
//type Number struct {
//	buf   []rune
//	lexer *Lexer
//}
//
//func (l *Number) Next(r rune) bool {
//	if unicode.IsDigit(r) {
//		l.buf = append(l.buf, r)
//		return true
//	}
//	l.lexer.NewToken(TNumber, l.buf)
//	l.lexer.SetState(StateMain)
//	l.buf = l.buf[:0]
//	return false
//}
//
//// Main — состояние по умолчанию.
//type Main struct {
//	lexer *Lexer
//}
//
//func (l *Main) Next(r rune) bool {
//	if unicode.IsDigit(r) {
//		l.lexer.SetState(StateNumber)
//		return false
//	}
//	return true
//}
//
//// Lexer содержит список состояний и полученные токены.
//type Lexer struct {
//	states []State
//	state  State
//	tokens []Token
//}
//
//// SetState изменяет состояние.
//func (lex *Lexer) SetState(state int) {
//	if state >= len(lex.states) {
//		panic("unknown state")
//	}
//	lex.state = lex.states[state]
//}
//
//// NewToken добавляет токен.
//func (lex *Lexer) NewToken(t TokenType, value []rune) {
//	lex.tokens = append(lex.tokens, Token{
//		Type:  t,
//		Value: string(value),
//	})
//}
//
//func main() {
//	var lex Lexer
//
//	// определяем состояния
//	lex.states = []State{&Main{lexer: &lex}, &Number{lexer: &lex}}
//	lex.SetState(StateMain)
//
//	// пробуем разобрать эту строку
//	s := "line778, 5 + 35 равно 40"
//	for _, ch := range s {
//		for !lex.state.Next(ch) {
//		}
//	}
//	// завершаем разбор последнего токена, если он начат
//	lex.state.Next(0)
//
//	fmt.Println(lex.tokens)
//}

//type Person struct {
//	Name string
//	Age  int
//}
//
//func (p Person) String() string {
//	return fmt.Sprintf("%s: %d", p.Name, p.Age)
//}
//
//type ByAge []Person
//
//// реализуем интерфейс sort.Interface для сортировки по возрасту
//
//func (a ByAge) Len() int           { return len(a) }
//func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
//func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
//
//// Sort сортирует слайс ByAge, так он реализует интерфейс sort.Interface.
//func (a ByAge) Sort() {
//	sort.Slice(a, func(i, j int) bool { return a[i].Name < a[j].Name })
//	sort.Slice(a, func(i, j int) bool { return a[i].Age < a[j].Age })
//}
//
//func main() {
//	people := ByAge{
//		{"Bob", 31},
//		{"John", 48}, // John старший
//		{"Michael", 17},
//		{"John", 26}, // John младший
//	}
//
//	fmt.Println(people)
//	// можем применить шаблонный метод
//	people.Sort()
//	fmt.Println(people)
//}
//
//var Funcs = make(map[string]func(r *http.Request)) // пустой интерфейс может принять любое значение
//
//func DBInsert(r *http.Request) {
//	// логика вставки
//}
//
//func DBDelete(r *http.Request) {
//	// логика удаления
//}
//
//func main() {
//	Funcs["DBInsert"] = DBInsert
//	Funcs["DBDelete"] = DBDelete
//	Funcs["DBChange"] = func(r *http.Request) {
//		// логика изменения
//	}
//	// ...
//}

//package main
//import (
//"fmt"
//"reflect"
//"unsafe"
//)
//func main() {
//	array := [10]byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
//	slice := array[:5]
//	// ограничили длину len слайса, но его вместимость cap не изменилась
//	fmt.Println(cap(slice)) // 10
//	// получим заголовок слайса
//	slHeader := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
//	// и в нём ограничим capacity
//	slHeader.Cap = 5
//	// так работает
//	fmt.Println(cap(slice)) // 5
//	// теперь, поскольку вместительность слайса ограничена,
//	// при вызове append ему будет выделен новый участок памяти
//	slice = append(slice, 42)
//	fmt.Println(slice) // [0 1 2 3 4 42]
//	// защитим несущий массив array от неожидаемых изменений
//	fmt.Println(array) // [0 1 2 3 4 5 6 7 8 9]
//}
//
//func ConditionalWrite(w io.Writer) (int, error) {
//	// используем рефлексию для определения конкретного типа
//	switch w.(type) {
//	case *os.File:
//		fmt.Print("File")
//	default:
//		fmt.Println("AAasddsdsdasd")
//		// если это *os.File, предпринимаем действия
//	}
//	// действуем по-другому
//	return 0, nil
//}

//type Storage map[string]string
//
//func (s Storage) Get(key string) (string, error) {
//	if _, ok := s[key]; !ok {
//		return "", fmt.Errorf("no key %s in map", key)
//	}
//	return s[key], nil
//}
