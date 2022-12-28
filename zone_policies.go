package acetime


// This is a sample zone_policies.go created by hand to help with developing the
// code that will parse and utilize these data structures. It will eventually be
// programmatically generated.

//---------------------------------------------------------------------------
// Policy name: US
// Rules: 5
// Memory (8-bit): 51
// Memory (32-bit): 72
//---------------------------------------------------------------------------

var ZoneRulesUS = []ZoneRule{
  // Rule    US    1967    2006    -    Oct    lastSun    2:00    0    S
  {
    1967 /*from_year*/,
    2006 /*to_year*/,
    10 /*in_month*/,
    7 /*on_day_of_week*/,
    0 /*on_day_of_month*/,
    8 /*at_time_code*/,
    0 /*at_time_modifier (kAtcSuffixW + minute=0)*/,
    4 /*delta_code ((delta_minutes=0)/15 + 4)*/,
    'S' /*letter*/,
  },
  // Rule    US    1976    1986    -    Apr    lastSun    2:00    1:00    D
  {
    1976 /*from_year*/,
    1986 /*to_year*/,
    4 /*in_month*/,
    7 /*on_day_of_week*/,
    0 /*on_day_of_month*/,
    8 /*at_time_code*/,
    0 /*at_time_modifier (kAtcSuffixW + minute=0)*/,
    8 /*delta_code ((delta_minutes=60)/15 + 4)*/,
    'D' /*letter*/,
  },
  // Rule    US    1987    2006    -    Apr    Sun>=1    2:00    1:00    D
  {
    1987 /*from_year*/,
    2006 /*to_year*/,
    4 /*in_month*/,
    7 /*on_day_of_week*/,
    1 /*on_day_of_month*/,
    8 /*at_time_code*/,
    0 /*at_time_modifier (kAtcSuffixW + minute=0)*/,
    8 /*delta_code ((delta_minutes=60)/15 + 4)*/,
    'D' /*letter*/,
  },
  // Rule    US    2007    max    -    Mar    Sun>=8    2:00    1:00    D
  {
    2007 /*from_year*/,
    9999 /*to_year*/,
    3 /*in_month*/,
    7 /*on_day_of_week*/,
    8 /*on_day_of_month*/,
    8 /*at_time_code*/,
    0 /*at_time_modifier (kAtcSuffixW + minute=0)*/,
    8 /*delta_code ((delta_minutes=60)/15 + 4)*/,
    'D' /*letter*/,
  },
  // Rule    US    2007    max    -    Nov    Sun>=1    2:00    0    S
  {
    2007 /*from_year*/,
    9999 /*to_year*/,
    11 /*in_month*/,
    7 /*on_day_of_week*/,
    1 /*on_day_of_month*/,
    8 /*at_time_code*/,
    0 /*at_time_modifier (kAtcSuffixW + minute=0)*/,
    4 /*delta_code ((delta_minutes=0)/15 + 4)*/,
    'S' /*letter*/,
  },

};

var ZonePolicyUS = ZonePolicy{
  ZoneRulesUS /*rules*/,
  nil /*letters*/,
};
