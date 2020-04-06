## XML

xml - генератор основы проекта: mfd файла, неймспейсов и сущностей в них. В качестве источника данных используется база данных. Результат - несколько xml файлов. 

### Использование

Генератор подключается к базе данных, считывает информацию о таблицах, отношениях между ними и генерирует xml на основе пользовательского ввода.
Если проект уже существует сначала генератор его загрузит и будет использовать как основу для будуших xml. Так же при вводе неймспейсов для таблиц будет предлагаться выбрать из существующих в проекте.

В результате выбранные таблицы сгенерируют сущность (entity). В сущности будут описаны все поля таблицы в виде атрибутов (attribute). Так же будут автоматически обавлены стандартные поиски (search). Сущности будут объеденены в неймспейсы и записаны в файлы с именем неймспейса.
Все неймспейсы будут записаны в mfd файл, который сохранится по указанному в команде пути.

### CLI
```
mfd-generator xml -h 

Create or update project base with namespaces and entities

Usage:
  mfd xml [flags]

Flags:
  -v, --verbose          print sql queries
  -c, --conn string      connection string to postgres database, e.g. postgres://usr:pwd@localhost:5432/db
  -m, --mfd string       mfd file path
  -t, --tables strings   table names for model generation separated by comma
                         use 'schema_name.*' to generate model for every table in model (default [public.*])
  -p, --nss string       use this parameter to set table & namespace in format "users=users,projects;shop=orders,prices"
  -h, --help             help for xml
```
  
`-t, --tables` - позволяет вводить исходные таблицы для генератора через запятую, если не указана схема для таблицы, то будет использоваться public. 
`*` - для генерирования всех таблиц в схеме, например: `public.*,geo.locations,geo.cities`      
`-n, --nss` - сайлент-режим, позволяет задать ассоциацию неймспейс - таблица.  
 
### MFD файл

Основоной файл проекта, содержит в себе настройки и список неймспейсов. Генериуется с нуля или дополняется.

```xml
<Project xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <Name>example.mfd</Name> <!-- имя проекта -->
    <PackageNames> <!-- список неймспейсов -->
        <string>blog</string>
        <string>common</string>
    </PackageNames>
    <Languages> <!-- список языков, см. генераторы xml-lang и template -->
        <string>en</string>
    </Languages> 
    <GoPGVer>8</GoPGVer> <!-- версия go-pg -->
</Project>
```

**PackageNames** - Указанные неймспейсы будут использоваться для дальнейшей генерации. Если неймспейс не указан в списке, даже если файл с неймспейсом присутствует, генерироваться не будет
**Languages** управление этим полем происходит в генераторе [xml-lang]. В дальнейшем генератор [template] будет использовать этот список чтобы сгенерировать языковые файлы для интерфейса vt      
**GoPGVer** - версия go-pg. Поддерживаемые значения 8 и 9. От этого параметра зависят все генераторы golang кода:
- импорты (`"github.com/go-pg/pg"` vs `"github.com/go-pg/pg/v9"`)
- аннотации к структурам (`sql:"title"` vs `pg:"title"`)
- функции (`pg.F` и `pg.Q` vs `pg.Ident` и `pg.SafeQuery`)
 
#### Namespace файл и сущности

Файл с неймспейсом, содержит все входящие в него сущности. Сущности будут сгруппированы в файлы по неймспейсам и в дальнейшей генерации

```xml
<Package xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
    <Name>blog</Name> <!-- имя неймспейса -->
    <Entities> <!-- список сущностей -->
        <Entity Name="Post" Namespace="blog" Table="post"> <!-- сущность -->
            <Attributes> <!-- список атрибутов -->
                <Attribute Name="ID" DBName="postId" DBType="int4" GoType="int" PK="true" Nullable="Yes" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="Alias" DBName="alias" DBType="varchar" GoType="string" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="255"></Attribute>
                <Attribute Name="Title" DBName="title" DBType="varchar" GoType="string" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="255"></Attribute>
                <Attribute Name="Text" DBName="text" DBType="text" GoType="string" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="Views" DBName="views" DBType="int4" GoType="int" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="CreatedAt" DBName="createdAt" DBType="timestamp" GoType="time.Time" PK="false" Nullable="No" Addable="false" Updatable="false" Min="0" Max="0"></Attribute>
                <Attribute Name="UserID" DBName="userId" DBType="int4" GoType="int" PK="false" FK="User" Nullable="No" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="TagID" DBName="tagId" IsArray="true" DBType="int4" GoType="[]int" PK="false" Nullable="Yes" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="StatusID" DBName="statusId" DBType="int4" GoType="int" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
            </Attributes>
            <Searches> <!-- список поисков -->
                <Search Name="IDs" AttrName="ID" SearchType="SEARCHTYPE_ARRAY"></Search>
                <Search Name="NotID" AttrName="ID" SearchType="SEARCHTYPE_NOT_EQUALS"></Search>
                <Search Name="TitleILike" AttrName="Title" SearchType="SEARCHTYPE_ILIKE"></Search>
                <Search Name="TextILike" AttrName="Text" SearchType="SEARCHTYPE_ILIKE"></Search>
            </Searches>
        </Entity>
        <Entity Name="Tag" Namespace="blog" Table="tags">
            <Attributes>
                <Attribute Name="ID" DBName="tagId" DBType="int4" GoType="int" PK="true" Nullable="Yes" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="Alias" DBName="alias" DBType="varchar" GoType="string" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="255"></Attribute>
                <Attribute Name="Title" DBName="title" DBType="varchar" GoType="string" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="255"></Attribute>
                <Attribute Name="Weight" DBName="weight" DBType="float8" GoType="*float64" PK="false" Nullable="Yes" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
                <Attribute Name="StatusID" DBName="statusId" DBType="int4" GoType="int" PK="false" Nullable="No" Addable="true" Updatable="true" Min="0" Max="0"></Attribute>
            </Attributes>
            <Searches>
                <Search Name="IDs" AttrName="ID" SearchType="SEARCHTYPE_ARRAY"></Search>
                <Search Name="NotID" AttrName="ID" SearchType="SEARCHTYPE_NOT_EQUALS"></Search>
                <Search Name="TitleILike" AttrName="Title" SearchType="SEARCHTYPE_ILIKE"></Search>
            </Searches>
        </Entity>
    </Entities>
</Package>
``` 

**Entity** - описание каждой сущности, содержит в себе имя (Name), неймспейс (Namespace) и соответствующую таблицу в бд (Table). 
В поле Table если не указана схема будет использоваться public  
Атрибут **Name** - содержит имя сущности, которое будет использоваться для генерирования имени структуры в генераторе [model]() и имени репозитория в генераторе [repo]() 
  
#### Атрибуты 

**Attribute** - содержит описание поля таблицы в бд, используется для генерирования vt aтрибутов в генераторе [xml-vt]() а та же для генерирования полей структуры в генераторе [model](model) и всех [зависимых](repo) [генераторах](vt)

**Name** - имя атрибута, используется для генерирования названия поля в структуре модели и поиска в генераторе [model](model). Уникально для сущности 
На это имя будут ссылаться поиски в разделе `<Searches>`, атрибуты `VTEntity` которые генерируются [xml-vt]()  
**DBName** - имя соответствующей колонки в таблице в бд.
**DBType** - тип соответствующей колонки в таблице в бд.
**GoType** - тип который будет сгенерирован генераторами golang кода. Поумолчанию для nullable атрибутов в типе будет присутствовать указатель, который можно убрать, если необходимо
**PK** - флаг Primary ключа, важное значение для генерирования vt-модели, функций типа GetByID, шаблонов и так далее
**Nullable** - Может ли значение быть nil. Флаг определяющий код модели и vt-модели (в конвертерах и валидаторах). Возможные значения `Yes` и `No`
**Addable** - Можно ли указать значение этого поля, при добавлении сущности в базу (например, ID). Флаг определяющий код модели vt-модели (в конвертерах). Возможные значения `true` и `false`  
**Updatable** - Можно ли указать значение этого поля, при обновлении сущности в базе (например, CreatedAt). Флаг определяющий код модели vt-модели (в конвертерах). Возможные значения `true` и `false` 
**Min** - Минимально возможное значение этого поля для чисел (например Age). Для строк - минимальное количество символов (например Description) Влияет на генерирование vt-модели (в валидаторах) 
**Max** - Максимально возможное значение этого поля (например Age). Для строк - максимальное количество символов (например Title) Влияет на генерирование vt-модели (в валидаторах)

#### Поиски
 
**Searches** - содержит в себе список полей для поиска по сущностям.

**Name** - Имя поиска в структуре поиска Search. Уникально для сущности, включая атрибуты.
**AttrName** - Ссылка на трибут сущности. Может быть ссылкой на другую сущность, в формате Entity.Attribute, например `User.ID` или `Category.ShowOnMain`. Влияет на генерирование всех моделей и [xml-vt]()
**SearchType** - Тип поиска, влияет на соотвествеющий тип поиска при постороении запросов в БД. Влияет на структуру Search и тип поля [возможные значения]()

Если атрибут добавляемый в модель новый (новая колонка в базе, новая таблица, новый проект) - то для этого атрибута будут сгенерированы поиски. 
Для строковых атрибутов - ilike, кроме поля Alias. Для ID - поиск по массиву IDs. Если присутсвует поле Alias, то добавляется поиск notID для генерирования поиска в vt- модели при проверке уникальности.

#### SEARCH_TYPE

Ниже приведены значения для поля SearchType и соответвующие им sql условия.
```
SEARCHTYPE_EQUALS         -  f = v          
SEARCHTYPE_NOT_EQUALS     -  f != v             
SEARCHTYPE_NULL           -  f is null       
SEARCHTYPE_NOT_NULL       -  f is not null            
SEARCHTYPE_GE             -  f >= v     
SEARCHTYPE_LE             -  f <= v      
SEARCHTYPE_G              -  f > v     
SEARCHTYPE_L              -  f < v     
SEARCHTYPE_LEFT_LIKE      -  f like '%v'             
SEARCHTYPE_LEFT_ILIKE     -  f ilike '%v'              
SEARCHTYPE_RIGHT_LIKE     -  f like 'v%'              
SEARCHTYPE_RIGHT_ILIKE    -  f ilike 'v%'               
SEARCHTYPE_LIKE           -  f like '%v%'        
SEARCHTYPE_ILIKE          -  f ilike '%v%'         
SEARCHTYPE_ARRAY          -  f in (v, v1, v2)         
SEARCHTYPE_NOT_INARRAY    -  f not in (v1, v2)               
SEARCHTYPE_ARRAY_CONTAIN  -  v any (f)                 
``` 
f - имя поля, v - значение
 

### Особенности проверки консистентности

При загрузке существующего проекта будет проведена проверка консистентности (это справедливо для всех генераторов).
Проверка включает в себя:
- каждый поиск в секции `<Searches>` ссылается на существующие в xml сущность и атрибут.
- каждый FK атрибут ссылается на а существующие в xml сущность и атрибут. 
В случае если проверки не пройдены - проект не загрузится с ошибкой.   
 
### Особенности работы с существующими сущностями

При повторной генерации генератор пытается сохранить пользовательские изменения:
- Если сущность для генерируемой таблицы уже существует, то она будет дополнена новыми атрибутами
- Новые атрибуты определяются по паре `DBName` и `DBType`. Это значит поля, у которых имя и тип уже пристуствуют в xml, добавляться не будут.
  - Если поменять тип колонки в таблице, то она будет добавлена в сущность как новая
- Если атрибут уже существует в xml, для него не будут сгенерированы новые поиски, даже если должны.
  - Если удалить поиск из секции `<Searches>` от при повторной генерации он не будет добавлен.  