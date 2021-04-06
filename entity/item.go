package entity

import "fmt"

type Element struct {
	Name string
}

type ElementsItem struct {
	Name           string
	ElementRecipes [3]*Element
}

func (e1 *ElementsItem) Conflict(e2 *ElementsItem) bool {
	if e1.Name == e2.Name {
		return true
	}

	for _, er1 := range e1.ElementRecipes {
		for _, er2 := range e2.ElementRecipes {
			if er1.Name == er2.Name {
				return true
			}
		}
	}

	return false
}

type Item struct {
	Name string
}

type UpgradedElementsItem struct {
	Name               string
	ElementsItemRecipe *ElementsItem
	ItemRecipe         *Item
}

func (u *UpgradedElementsItem) ContainsElementsItem(e *ElementsItem) bool {
	if u.ElementsItemRecipe.Name == e.Name {
		return true
	}

	return false
}

func (u *UpgradedElementsItem) Print() {
	fmt.Println(u.ElementsItemRecipe.Name, "/", u.ItemRecipe.Name)
}

var (
	Ice    = &Element{"Ice"}
	Fire   = &Element{"Fire"}
	Water  = &Element{"Water"}
	Energy = &Element{"Energy"}
	Earth  = &Element{"Earth"}
	Life   = &Element{"Life"}
	Void   = &Element{"Void"}
	Air    = &Element{"Air"}
	Light  = &Element{"Light"}
	Shadow = &Element{"Shadow"}
)

var (
	FireSword         = &ElementsItem{"Fire Sword", [3]*Element{Fire, Earth, Light}}
	SpearOfLife       = &ElementsItem{"Spear of Life", [3]*Element{Earth, Life, Light}}
	EarthSpear        = &ElementsItem{"Earth Spear", [3]*Element{Earth, Air, Shadow}}
	StoneOfVoid       = &ElementsItem{"Stone of Void", [3]*Element{Life, Void, Air}}
	FireStone         = &ElementsItem{"Fire Stone", [3]*Element{Fire, Life, Air}}
	VoidTalisman      = &ElementsItem{"Void Talisman", [3]*Element{Fire, Energy, Void}}
	RingOfWater       = &ElementsItem{"Ring of Water", [3]*Element{Water, Energy, Shadow}}
	EnergyShield      = &ElementsItem{"Energy Shield", [3]*Element{Water, Energy, Earth}}
	FireHammer        = &ElementsItem{"Fire Hammer", [3]*Element{Fire, Water, Earth}}
	RingOfIce         = &ElementsItem{"Ring of Ice", [3]*Element{Ice, Air, Shadow}}
	EnergyAxe         = &ElementsItem{"Energy Axe", [3]*Element{Ice, Energy, Void}}
	IceSword          = &ElementsItem{"Ice Sword", [3]*Element{Ice, Light, Shadow}}
	WaterBlades       = &ElementsItem{"Water Blades", [3]*Element{Water, Life, Shadow}}
	StaffOfLight      = &ElementsItem{"Staff of Light", [3]*Element{Water, Void, Light}}
	AmuletOfVampirism = &ElementsItem{"Amulet of Vampirism", [3]*Element{Ice, Earth, Void}}
	ShadowSword       = &ElementsItem{"Shadow Sword", [3]*Element{Fire, Void, Shadow}}
	ShieldOfLight     = &ElementsItem{"Shield of Light", [3]*Element{Energy, Air, Light}}
	LifeStone         = &ElementsItem{"Life Stone", [3]*Element{Ice, Life, Light}}
	FireBlades        = &ElementsItem{"Fire Blades", [3]*Element{Fire, Energy, Life}}
	IceShield         = &ElementsItem{"Ice Shield", [3]*Element{Ice, Water, Air}}
	WaterStone        = &ElementsItem{"Water Stone ", [3]*Element{Water, Void, Air}}
	MysteryStone      = &ElementsItem{"Mystery Stone", [3]*Element{Water, Void, Shadow}}
	EnergyBlade       = &ElementsItem{"Energy Blade", [3]*Element{Energy, Light, Shadow}}
	LightCrossbow     = &ElementsItem{"Light Crossbow", [3]*Element{Fire, Air, Light}}
	FireStaff         = &ElementsItem{"Fire Staff", [3]*Element{Ice, Fire, Water}}
	IceBlade          = &ElementsItem{"Ice Blade", [3]*Element{Ice, Fire, Earth}}
	EnergyBow         = &ElementsItem{"Energy Bow", [3]*Element{Energy, Earth, Air}}
	UrnOfLife         = &ElementsItem{"Urn of Life", [3]*Element{Life, Void, Light}}
	LifeShield        = &ElementsItem{"Life Shield", [3]*Element{Earth, Life, Shadow}}
	IceShard          = &ElementsItem{"Ice Shard", [3]*Element{Ice, Energy, Life}}
)

var AllElementsItems = []*ElementsItem{FireSword, SpearOfLife, EarthSpear, StoneOfVoid, FireStone, VoidTalisman, RingOfWater, EnergyShield, FireHammer, RingOfIce, EnergyAxe, IceSword, WaterBlades, StaffOfLight, AmuletOfVampirism, ShadowSword, ShieldOfLight, LifeStone, FireBlades, IceShield, WaterStone, MysteryStone, EnergyBlade, LightCrossbow, FireStaff, IceBlade, EnergyBow, UrnOfLife, LifeShield, IceShard}

var (
	IceEssence    = &ElementsItem{"Ice Essence", [3]*Element{Ice, Ice, Ice}}
	FireEssence   = &ElementsItem{"Fire Essence", [3]*Element{Fire, Fire, Fire}}
	WaterEssence  = &ElementsItem{"Water Essence", [3]*Element{Water, Water, Water}}
	EnergyEssence = &ElementsItem{"Energy Essence", [3]*Element{Energy, Energy, Energy}}
	EarthEssence  = &ElementsItem{"Earth Essence", [3]*Element{Earth, Earth, Earth}}
	LifeEssence   = &ElementsItem{"Life Essence", [3]*Element{Life, Life, Life}}
	VoidEssence   = &ElementsItem{"Void Essence", [3]*Element{Void, Void, Void}}
	AirEssence    = &ElementsItem{"Air Essence", [3]*Element{Air, Air, Air}}
	LightEssence  = &ElementsItem{"Light Essence", [3]*Element{Light, Light, Light}}
	ShadowEssence = &ElementsItem{"Shadow Essence", [3]*Element{Shadow, Shadow, Shadow}}
)

var (
	Radiance              = &Item{"Radiance"}
	GuardianGreaves       = &Item{"Guardian Greaves"}
	AssaultCuirass        = &Item{"Assault Cuirass"}
	SolarCrest            = &Item{"Solar Crest"}
	VladmirsOffering      = &Item{"Vladmir's Offering"}
	RodOfAtos             = &Item{"Rod of Atos"}
	Butterfly             = &Item{"Butterfly"}
	LinkensSphere         = &Item{"Linken's Sphere"}
	SangeandYasha         = &Item{"Sange and Yasha"}
	PipeOfInsight         = &Item{"Pipe of Insight"}
	OctarineCore          = &Item{"Octarine Core"}
	Dagon5                = &Item{"Dagon 5"}
	BladeMail             = &Item{"Blade Mail"}
	DragonLance           = &Item{"Dragon Lance"}
	Satanic               = &Item{"Satanic"}
	EchoSabre             = &Item{"Echo Sabre"}
	HeartOfTarrasque      = &Item{"Heart of Tarrasque"}
	Daedalus              = &Item{"Daedalus"}
	Desolator             = &Item{"Desolator"}
	ShivasGuard           = &Item{"Shiva's Guard"}
	DiffusalBlade         = &Item{"Diffusal Blade"}
	EulsScepterOfDivinity = &Item{"Eul's Scepter of Divinity"}
	VeilOfDiscord         = &Item{"Veil of Discord"}
	Mjolnir               = &Item{"Mjolnir"}
	EyeOfSkadi            = &Item{"Eye of Skadi"}
	MonkeyKingBar         = &Item{"Monkey King Bar"}
	SpiritVessel          = &Item{"Spirit Vessel"}
	CrimsonGuard          = &Item{"Crimson Guard"}
	Aluneth               = &Item{"Aluneth"}
)

var (
	FireRadiance             = &UpgradedElementsItem{"Fire Radiance", FireSword, Radiance}
	LifeTreads               = &UpgradedElementsItem{"Life Treads", SpearOfLife, GuardianGreaves}
	EarthCuirass             = &UpgradedElementsItem{"Earth Cuirass", EarthSpear, AssaultCuirass}
	VoidShield               = &UpgradedElementsItem{"Void Shield", StoneOfVoid, SolarCrest}
	VampireRobe              = &UpgradedElementsItem{"Vampire Robe", FireStone, VladmirsOffering}
	TalismanOfAtos           = &UpgradedElementsItem{"Talisman of Atos", VoidTalisman, RodOfAtos}
	WaterButterfly           = &UpgradedElementsItem{"Water Butterfly", RingOfWater, Butterfly}
	EnergySphere             = &UpgradedElementsItem{"Energy Sphere", EnergyShield, LinkensSphere}
	DemonicSword             = &UpgradedElementsItem{"Demonic Sword", FireHammer, SangeandYasha}
	PipeOfIce                = &UpgradedElementsItem{"Pipe of Ice", RingOfIce, PipeOfInsight}
	EnergyCore               = &UpgradedElementsItem{"Energy Core", EnergyAxe, OctarineCore}
	IceDagon                 = &UpgradedElementsItem{"Ice Dagon", IceSword, Dagon5}
	WaterCarapace            = &UpgradedElementsItem{"Water Carapace", WaterBlades, BladeMail}
	StaffOfDragon            = &UpgradedElementsItem{"Staff of Dragon", StaffOfLight, DragonLance}
	UnhallowedIcon           = &UpgradedElementsItem{"Unhallowed Icon", AmuletOfVampirism, Satanic}
	Kingsbane                = &UpgradedElementsItem{"Kingsbane", ShadowSword, EchoSabre}
	HeartOfLight             = &UpgradedElementsItem{"Heart of Light", ShieldOfLight, HeartOfTarrasque}
	TheCausticFinale         = &UpgradedElementsItem{"The Caustic Finale", LifeStone, Daedalus}
	FireDesolator            = &UpgradedElementsItem{"Fire Desolator", FireBlades, Desolator}
	ShivasShield             = &UpgradedElementsItem{"Shiva's Shield", IceShield, ShivasGuard}
	ManaBlade                = &UpgradedElementsItem{"Mana Blade", WaterStone, DiffusalBlade}
	MysteryScepterOfDivinity = &UpgradedElementsItem{"Mystery Scepter of Divinity", MysteryStone, EulsScepterOfDivinity}
	EnergyVeil               = &UpgradedElementsItem{"Energy Veil", EnergyBlade, VeilOfDiscord}
	HammerOfGod              = &UpgradedElementsItem{"Hammer of God", LightCrossbow, Mjolnir}
	FireCore                 = &UpgradedElementsItem{"Fire Core", FireStaff, OctarineCore}
	SkadiBow                 = &UpgradedElementsItem{"Skadi Bow", IceBlade, EyeOfSkadi}
	MonkeyKingBow            = &UpgradedElementsItem{"Monkey King Bow", EnergyBow, MonkeyKingBar}
	LifeVessel               = &UpgradedElementsItem{"Life Vessel", UrnOfLife, SpiritVessel}
	CuirassOfLife            = &UpgradedElementsItem{"Cuirass of Life", LifeShield, CrimsonGuard}
	IceAluneth               = &UpgradedElementsItem{"Ice Aluneth", IceShard, Aluneth}
)

var AllUpgradedElementsItems = []*UpgradedElementsItem{FireRadiance, LifeTreads, EarthCuirass, VoidShield, VampireRobe, TalismanOfAtos, WaterButterfly, EnergySphere, DemonicSword, PipeOfIce, EnergyCore, IceDagon, WaterCarapace, StaffOfDragon, UnhallowedIcon, Kingsbane, HeartOfLight, TheCausticFinale, FireDesolator, ShivasShield, ManaBlade, MysteryScepterOfDivinity, EnergyVeil, HammerOfGod, FireCore, SkadiBow, MonkeyKingBow, LifeVessel, CuirassOfLife, IceAluneth}

var (
	UpgradedButterfly     = &UpgradedElementsItem{"Upgraded Butterfly", IceEssence, Butterfly}
	UpgradedDaedalus      = &UpgradedElementsItem{"Upgraded Daedalus", FireEssence, Daedalus}
	UpgradedSangeandYasha = &UpgradedElementsItem{"Upgraded Sange and Yasha", WaterEssence, SangeandYasha}
	UpgradedMjolnir       = &UpgradedElementsItem{"Upgraded Mjolnir", EnergyEssence, Mjolnir}
	UpgradedHeart         = &UpgradedElementsItem{"Upgraded Heart", EarthEssence, HeartOfTarrasque}
	UpgradedSatanic       = &UpgradedElementsItem{"Upgraded Satanic", LifeEssence, Satanic}
	UpgradedDiffusalBlade = &UpgradedElementsItem{"Upgraded Diffusal Blade", VoidEssence, DiffusalBlade}
	UpgradedPipeOfInsight = &UpgradedElementsItem{"Upgraded Pipe of Insight", AirEssence, PipeOfInsight}
	UpgradedMonkeyKingBar = &UpgradedElementsItem{"Upgraded Monkey King Bar", LightEssence, MonkeyKingBar}
	UpgradedDesolator     = &UpgradedElementsItem{"Upgraded Desolator", ShadowEssence, Desolator}
)
